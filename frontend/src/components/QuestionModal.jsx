import React, { useState } from 'react';
import './QuestionModal.css'; // We will create this

const QuestionModal = ({ question, onAnswer }) => {
    if (!question) return null;

    const [selectedOption, setSelectedOption] = useState(null);

    const handleSubmit = () => {
        if (selectedOption) {
            onAnswer(selectedOption);
            setSelectedOption(null); // Reset
        }
    };

    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <h3>Question Time!</h3>
                <p className="question-text">{question.text}</p>
                <div className="options-grid">
                    {question.options.map((option, index) => (
                        <button 
                            key={index} 
                            className={`option-btn ${selectedOption === option ? 'selected' : ''}`}
                            onClick={() => setSelectedOption(option)}
                        >
                            {option}
                        </button>
                    ))}
                </div>
                <div className="modal-actions">
                    <button 
                        className="submit-btn" 
                        onClick={handleSubmit} 
                        disabled={!selectedOption}
                    >
                        Submit Answer
                    </button>
                </div>
                <div className="difficulty-badge">
                    Difficulty: {question.difficulty}
                </div>
            </div>
        </div>
    );
};

export default QuestionModal;
