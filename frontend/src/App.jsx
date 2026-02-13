
import React, { useEffect } from 'react';
import { useGame } from './hooks/useGame';
import Board from './components/Board';
// import QuestionModal from './components/QuestionModal'; // Removed
import GameStatus from './components/GameStatus';

function App() {
    const { 
        gameState, 
        loading, 
        error, 
        initGame, 
        // handleAnswer, // Removed
        // modalOpen,    // Removed
        // currentQuestion, // Removed
        message
    } = useGame();

    useEffect(() => {
        // START NEW GAME ON LOAD
        const startNewGame = () => {
            // Calculate grid size based on FULL screen
            const CELL_SIZE = 40; 
            const width = window.innerWidth;
            const height = window.innerHeight;
            
            const cols = Math.floor(width / CELL_SIZE);
            const rows = Math.floor(height / CELL_SIZE);
            
            // Ensure odd numbers for better maze generation
            const finalCols = cols % 2 === 0 ? cols - 1 : cols;
            const finalRows = rows % 2 === 0 ? rows - 1 : rows;

            initGame(finalRows, finalCols);
        };

        startNewGame();
        
        // Add resize listener? For now, we only generate on start.
    }, [initGame]);

    if (loading && !gameState) return <div className="loading-screen"><h2>Generating Huge Maze...</h2></div>;
    if (error) return <div className="error-screen"><h2 style={{color: 'red'}}>Error: {error}</h2><button onClick={() => window.location.reload()}>Retry</button></div>;

    return (
        <div className="game-container">
            {/* HUD Overlay */}
            <div className="hud-overlay">
                <div className="hud-header">
                    <h1>GO Maze Game</h1>
                    <p>Escape the Labyrinth. Find the Hidden Exit.</p>
                </div>
                
                {gameState && (
                    <div className="hud-status">
                         {/* Optional: Keep status if we track time or steps later? For now keeping it for lives/score if applicable or removing it? 
                             User said "remove questions". Score usually comes from questions. 
                             Lives come from wrong answers. 
                             So maybe GameStatus is useless now? 
                             Let's keep it to show "Player Ready" or remove it. 
                             I'll keep it but maybe it looks empty.
                         */}
                        <GameStatus player={gameState.player} message={message} />
                    </div>
                )}
            </div>

            {/* Main Game Board */}
            {gameState && (
                <div className="board-wrapper">
                    <Board 
                        board={gameState.board} 
                        playerPos={gameState.player.current_pos} 
                    />

                    {gameState.status === 'LOST' && (
                        <div className="game-over-overlay">
                             <h2 style={{color: 'red'}}>GAME OVER</h2>
                             <button onClick={() => window.location.reload()}>Try Again</button>
                        </div>
                    )}

                    {gameState.status === 'WON' && ( 
                        <div className="game-over-overlay">
                             <h2 style={{color: 'green'}}>VICTORY!</h2>
                             <button onClick={() => window.location.reload()}>Play Again</button>
                        </div>
                    )}
                </div>
            )}
        </div>
    );
}

export default App;
