import React from 'react';
import Cell from './Cell';
import './Board.css';

const Board = ({ board, playerPos }) => {
    if (!board || !board.grid) return <div className="loading">Loading Board...</div>;

    const VISIBILITY_RADIUS = 3; // How far player can see

    return (
        <div className="board-container">
            <div 
                className="board"
                style={{
                    gridTemplateColumns: `repeat(${board.cols}, 40px)`,
                    gridTemplateRows: `repeat(${board.rows}, 40px)`
                }}
            >
                {board.grid.map((row, rowIndex) => (
                    row.map((cell, colIndex) => {
                        // Calculate Distance
                        const dx = Math.abs(playerPos.x - colIndex);
                        const dy = Math.abs(playerPos.y - rowIndex);
                        const dist = Math.sqrt(dx*dx + dy*dy);
                        
                        const isVisible = dist <= VISIBILITY_RADIUS;
                        
                        // If hidden, render Fog
                        if (!isVisible) {
                             return <div key={`${rowIndex}-${colIndex}`} className="cell fog"></div>;
                        }
                        
                        // If visible, render actual cell
                        const isPlayerHere = playerPos.x === colIndex && playerPos.y === rowIndex;
                        return (
                            <Cell 
                                key={`${rowIndex}-${colIndex}`} 
                                cell={cell} 
                                isPlayerHere={isPlayerHere} 
                            />
                        );
                    })
                ))}
            </div>
        </div>
    );
};

export default Board;
