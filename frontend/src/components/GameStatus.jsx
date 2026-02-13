import React from 'react';

const GameStatus = ({ player, message }) => {
    if (!player) return null;

    return (
        <div style={{
            display: 'flex', 
            justifyContent: 'space-between', 
            alignItems: 'center',
            background: '#161b22',
            padding: '10px 20px',
            borderRadius: '8px',
            marginBottom: '10px',
            borderBottom: '2px solid var(--primary-color)'
        }}>
            <div style={{ color: 'var(--danger-color)', fontSize: '1.2rem' }}>
                Lives: {'❤️'.repeat(player.lives)}
            </div>
            
            <div style={{ fontSize: '1.1rem', fontWeight: 'bold' }}>
                {message && <span style={{ color: 'var(--accent-color)', marginRight: '20px' }}>{message}</span>}
            </div>

            <div style={{ color: 'var(--primary-color)', fontSize: '1.2rem' }}>
                Score: {player.score}
            </div>
        </div>
    );
};

export default GameStatus;
