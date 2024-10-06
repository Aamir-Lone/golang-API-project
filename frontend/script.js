const apiUrl = "http://localhost:4000/api/todos"; // Backend API URL
const todoList = document.getElementById("todo-list");
const form = document.getElementById("todo-form");

// Fetch and display all todos
async function fetchTodos() {
    const response = await fetch(apiUrl);
    const todos = await response.json();
    todoList.innerHTML = "";
    todos.forEach(todo => displayTodo(todo));
}

// Display a single todo
function displayTodo(todo) {
    const todoElement = document.createElement("div");
    todoElement.classList.add("todo");
    todoElement.style.backgroundColor = todo.completed ? "#d4edda" : "#f8d7da"; // Green if completed, red if not

    todoElement.innerHTML = `
        <h3>${todo.title}</h3>
        <p>Description: ${todo.description || "No description"}</p>
        <p>Date Created: ${new Date(todo.dateCreated).toLocaleDateString()}</p>
        <p>Completed: ${todo.completed ? "Yes" : "No"}</p>
        <button onclick="deleteTodo('${todo._id}')">Delete</button>
        <button onclick="toggleComplete('${todo._id}', ${todo.completed})">Toggle Complete</button>
    `;
    todoList.appendChild(todoElement);
}

// Add a new todo
form.addEventListener("submit", async (event) => {
    event.preventDefault();
    const title = document.getElementById("title").value;
    const description = document.getElementById("description").value;
    const completed = false;
    const dateCreated = new Date().toISOString(); // Automatically set the current date

    const newTodo = { title, description, completed, dateCreated };
    await fetch(apiUrl, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(newTodo)
    });

    form.reset();
    fetchTodos();
});

// Delete a todo
async function deleteTodo(id) {
    console.log(`Deleting todo with id: ${id}`);
    const response = await fetch(`http://localhost:4000/api/todo/${id}`, { method: "DELETE" }); // Changed to use path parameter
    console.log(response);
    if (response.ok) {
        console.log(`Todo with id ${id} deleted successfully`);
        fetchTodos(); // Refresh the list
    } else {
        console.error("Failed to delete todo");
    }
}

// Toggle the completion status
async function toggleComplete(id, currentStatus) {
    console.log(`Toggling completion for todo with id: ${id}`);
    const updatedTodo = { completed: !currentStatus };
    const response = await fetch(`http://localhost:4000/api/todo/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(updatedTodo)
    });
    console.log(response);
    if (response.ok) {
        console.log(`Todo with id ${id} updated successfully`);
        fetchTodos(); // Refresh the list
    } else {
        console.error("Failed to update todo");
    }
}

// Initial fetch
fetchTodos();
