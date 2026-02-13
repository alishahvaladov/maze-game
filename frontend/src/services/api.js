const API_BASE_URL = 'http://localhost:3000/api';

export const startGame = async (rows, cols) => {
    try {
        const response = await fetch(`${API_BASE_URL}/game/start`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ rows, cols }),
        });
        if (!response.ok) throw new Error('Failed to start game');
        return await response.json();
    } catch (error) {
        console.error("API Error:", error);
        throw error;
    }
};

export const movePlayer = async (gameId, direction) => {
    try {
        const response = await fetch(`${API_BASE_URL}/game/${gameId}/move`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ direction }),
        });
        if (!response.ok) throw new Error('Failed to move');
        return await response.json();
    } catch (error) {
        console.error("API Error:", error);
        throw error;
    }
};

export const answerQuestion = async (gameId, questionId, answer) => {
    try {
        const response = await fetch(`${API_BASE_URL}/game/${gameId}/answer`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ question_id: questionId, answer }),
        });
        if (!response.ok) throw new Error('Failed to answer');
        return await response.json();
    } catch (error) {
        console.error("API Error:", error);
        throw error;
    }
};
