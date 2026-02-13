import React from 'react';
import './Cell.css';

const Cell = ({ cell, isPlayerHere }) => {
    // Determine base class
    let className = 'cell';
    
    if (cell.type === 'WALL') {
        className += ' wall';
        if (cell.is_question_wall) className += ' wall-question';
    } else if (cell.type === 'START') {
        className += ' path start';
    } else if (cell.type === 'EXIT') { // Backend returns EXIT
        className += ' path end';
    } else if (cell.type === 'BEYOND') {
        className += ' beyond';
    } else {
        className += ' path';
    }

    return (
        <div className={className} title={`Pos: ${cell.position.x},${cell.position.y}`}>
            {isPlayerHere && <div className="player-token" />}
            
            {/* Standard Path Question (?) */}
            {!isPlayerHere && cell.has_question && !cell.is_question_wall && (
                <div className="question-marker">?</div>
            )}
            
            {/* Question Wall Lock (Handled by CSS mostly, but strictly no '?' here) */}
            
            {cell.type === 'EXIT' && !isPlayerHere && <span>🏁</span>}
        </div>
    );
};

export default React.memo(Cell);
