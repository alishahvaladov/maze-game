# Go Maze Game - The Labyrinth

A fullscreen, web-based labyrinth exploration game built with **Golang** (Backend) and **React + Vite** (Frontend).

Navigate a massive, procedurally generated maze that fills your entire screen. Beware the **Fog of War**—you can only see a few steps ahead. Your goal is to find the hidden Exit.

---

## 🚀 Getting Started

### Prerequisites

*   **Go** (v1.20 or higher) - [Download Go](https://go.dev/dl/)
*   **Node.js** (v16 or higher) & **npm** (or yarn) - [Download Node.js](https://nodejs.org/)

### Installation & Setup

1.  **Clone the Repository** (if you haven't already):
    ```bash
    git clone https://github.com/yourusername/go-maze-game.git
    cd go-maze-game
    ```

2.  **Start the Backend (Game Server)**:
    Open a terminal in the project root:
    ```bash
    cd backend
    # Install dependencies (if any external ones are used, e.g. godotenv)
    go mod tidy
    # Run the server
    go run main.go
    ```
    The backend will start on `http://localhost:3000`.

3.  **Start the Frontend (Web Client)**:
    Open a **new** terminal window/tab:
    ```bash
    cd frontend
    # Install dependencies
    npm install
    # Start the development server
    npm run dev
    ```
    The frontend will usually start on `http://localhost:5173`. Open this URL in your browser.

---

## 🎮 How to Play

1.  **Objective**: You are the **Purple Square**. You must find the **Green Exit** hidden somewhere in the maze.
2.  **Controls**: Use your **Arrow Keys** (Up, Down, Left, Right) to move.
3.  **Fog of War**: The maze is shrouded in darkness. You can only see a small radius around your player. Explore to reveal the map!
4.  **The Maze**: It is a "Braided" maze, meaning there are loops and dead ends. You can get lost if you don't pay attention.
5.  **Winning**: Reach the Green Exit tile to trigger the Victory screen.

---

## 🛠️ Technology Stack

*   **Backend**: 
    *   Language: Go (Golang)
    *   Framework: Standard `net/http` library
    *   Algorithm: Recursive Backtracking (DFS) with Braiding for complex loops
*   **Frontend**:
    *   Framework: React.js
    *   Build Tool: Vite
    *   Styling: CSS Modules & Custom CSS (Dark Mode theme)
    *   State Management: React Hooks (`useGame`, `useState`, `useEffect`)

---

## 🐛 Troubleshooting

*   **"Maze is just a line"**:
    *   Make sure the Backend is running *before* you load the Frontend.
    *   Refresh the page. The game calculates grid size based on window dimensions on load.
*   **"Connection Refused"**:
    *   Ensure the backend is running on port 3000. Check terminal logs for errors.

---

## 📜 License

This project is open source and available under the [MIT License](LICENSE).
