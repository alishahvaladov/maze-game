import { useState, useEffect, useCallback, useRef } from 'react';
import { startGame } from '../services/api';

export const useGame = () => {
    // Game State
    const [gameState, setGameState] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    
    // UI State
    const [message, setMessage] = useState(""); 
    
    const socketRef = useRef(null);

    // Initial Start
    const initGame = useCallback(async (rows, cols) => {
        setLoading(true);
        setError(null);
        try {
            // 1. Start Game via HTTP to get ID and initial state
            const data = await startGame(rows, cols);
            setGameState(data);

            // 2. Connect WebSocket
            // Determine protocol (ws vs wss) and host
            const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
            // In dev, frontend is 5173, backend is 3000. 
            // If running strictly locally we might need to point to localhost:3000 explicitly if not proxied.
            // Assuming localhost:3000 for backend based on previous context.
            const host = "localhost:3000"; 
            const wsUrl = `${protocol}//${host}/ws?id=${data.id}`;
            
            console.log("Connecting to WS:", wsUrl);
            const socket = new WebSocket(wsUrl);

            socketRef.current = socket;

            socket.onopen = () => {
                console.log("WebSocket Connected");
            };

            socket.onmessage = (event) => {
                try {
                    const response = JSON.parse(event.data);
                    if (response.type === 'update') {
                         const payload = response.payload;
                         
                         // Update State
                         setGameState(prevState => {
                             if (!prevState) return null; // Should not happen if initGame ran
                             
                             // If full game_state is provided (e.g. game over or special event), use it.
                             if (payload.game_state) {
                                 return payload.game_state;
                             }
                             
                             // Otherwise merge partial updates (Player, Status)
                             return {
                                 ...prevState,
                                 player: payload.player || prevState.player,
                                 status: payload.status || prevState.status
                             };
                         });
                         
                         // Simple result handling
                         if (payload.result === "Blocked") {
                                setMessage("Path Blocked!");
                                setTimeout(() => setMessage(""), 500);
                         } else if (payload.result === "Win") {
                                setMessage("You Won!");
                         }
                    } else if (response.type === 'error') {
                        console.error("WS Error:", response.payload);
                    }
                } catch (e) {
                    console.error("WS Message Parse Error", e);
                }
            };

            socket.onclose = () => {
                console.log("WebSocket Disconnected");
            };
            
            socket.onerror = (err) => {
                console.error("WebSocket Error:", err);
            };

        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    }, []);

    // Handle Movement (WebSocket)
    const handleMove = useCallback((direction) => {
        if (!gameState || gameState.status !== 'ACTIVE') return;
        
        if (socketRef.current && socketRef.current.readyState === WebSocket.OPEN) {
            const msg = {
                type: "move",
                direction: direction
            };
            socketRef.current.send(JSON.stringify(msg));
        }
    }, [gameState]);

    // Keyboard Listeners
    useEffect(() => {
        const onKeyDown = (e) => {
            // Only capture if game is active
            if (!gameState) return;

            switch(e.key) {
                case "ArrowUp": handleMove("UP"); break;
                case "ArrowDown": handleMove("DOWN"); break;
                case "ArrowLeft": handleMove("LEFT"); break;
                case "ArrowRight": handleMove("RIGHT"); break;
                default: break;
            }
        };
        window.addEventListener('keydown', onKeyDown);
        return () => window.removeEventListener('keydown', onKeyDown);
    }, [handleMove, gameState]);

    // Cleanup on unmount
    useEffect(() => {
        return () => {
            if (socketRef.current) {
                socketRef.current.close();
            }
        };
    }, []);

    return {
        gameState,
        loading,
        error,
        initGame,
        handleMove,
        message
    };
};
