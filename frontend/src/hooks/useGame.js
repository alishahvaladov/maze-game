import { useState, useEffect, useCallback } from 'react';
import { startGame, movePlayer, answerQuestion } from '../services/api';

export const useGame = () => {
    // Game State
    const [gameState, setGameState] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    
    // UI State
    const [modalOpen, setModalOpen] = useState(false);
    const [currentQuestion, setCurrentQuestion] = useState(null);
    const [message, setMessage] = useState(""); // For small notifications like "Blocked"

    // Initial Start
    const initGame = useCallback(async (rows, cols) => {
        setLoading(true);
        setError(null);
        try {
            const data = await startGame(rows, cols);
            setGameState(data);
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    }, [/* No dependencies needed as startGame is imported and setters are stable */]);

    // Handle Movement
    const handleMove = useCallback(async (direction) => {
        if (!gameState || modalOpen || gameState.status !== 'ACTIVE') return;

        try {
            const result = await movePlayer(gameState.id, direction);
            
            // Log result for debugging
            console.log("Move result:", result);

            // Handle Move Result (Message from backend)
            // Expecting result to have: { status: "Moved"|"Blocked"|"QuestionFound"|"QuestionWallHit"|"Win", game_state: {...}, question: {...} }
            // Note: Backend 'MovePlayer' currently only returns string,error. 
            // We need to update backend to return a full JSON object in handlers.go
            // Assuming for now the backend handler will be updated to return { result: string, game_state: object, question: object }
            
            // Update state if provided
            if (result.game_state) {
                setGameState(result.game_state);
            }

            // Handle events
            switch (result.result) {
                case "Blocked":
                    setMessage("Path Blocked!");
                    setTimeout(() => setMessage(""), 1000);
                    break;
                case "QuestionFound":
                    // Trigger Question Modal
                    // Question data should be in response
                    if (result.question) {
                        setCurrentQuestion(result.question);
                        setModalOpen(true);
                    }
                    else {
                        // Fallback purely for dev if backend isn't perfect yet
                        console.warn("QuestionFound but no question data sent.");
                    }
                    break;
                case "QuestionWallHit":
                    // Trigger Hard Question
                    if (result.question) {
                        setCurrentQuestion(result.question);
                        setModalOpen(true);
                        setMessage("It's a Trap! Answer to pass.");
                    }
                    break;
                case "Win":
                    setMessage("You Won!");
                    break;
                default:
                    // Just moved
                    break;
            }

        } catch (err) {
            console.error("Move error:", err);
        }
    }, [gameState, modalOpen]);

    // Handle Answer
    const handleAnswer = async (answer) => {
        if (!gameState || !currentQuestion) return;

        try {
            const result = await answerQuestion(gameState.id, currentQuestion.id, answer);
            
            if (result.correct) {
                setMessage("Correct!");
                setTimeout(() => setMessage(""), 1000);
                setModalOpen(false);
                setCurrentQuestion(null);
                // Update state
                if (result.game_state) {
                    setGameState(result.game_state);
                }
            } else {
                setMessage("Wrong! Lost a life.");
                // Update state to show reduced life
                if (result.game_state) {
                    setGameState(result.game_state);
                }
                // Close modal if life > 0? Or keep trying? 
                // Usually maze games might just subtract life and keep you there or move you back.
                // For now, let's close modal.
                setModalOpen(false);
                setCurrentQuestion(null);
            }

        } catch (err) {
            console.error(err);
        }
    };

    // Keyboard Listeners
    useEffect(() => {
        const onKeyDown = (e) => {
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
    }, [handleMove]);

    return {
        gameState,
        loading,
        error,
        initGame,
        handleMove,
        handleAnswer,
        modalOpen,
        currentQuestion,
        message
    };
};
