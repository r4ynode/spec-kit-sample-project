import React, { useState, useEffect } from 'react';
import TaskList from './components/TaskList';
import AddTaskForm from './components/AddTaskForm';
import { getTasks, createTask, updateTask, deleteTask } from './services/api';
import './App.css';

function App() {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    try {
      const data = await getTasks();
      setTasks(data);
    } catch (error) {
      console.error("Error fetching tasks:", error);
    }
  };

  const handleAddTask = async (title) => {
    try {
      const newTask = await createTask(title);
      setTasks([...tasks, newTask]);
    } catch (error) {
      console.error("Error creating task:", error);
    }
  };

  const handleToggleComplete = async (id) => {
    const taskToUpdate = tasks.find(task => task.id === id);
    if (taskToUpdate) {
      const updatedTask = { ...taskToUpdate, completed: !taskToUpdate.completed };
      try {
        await updateTask(id, { title: updatedTask.title, completed: updatedTask.completed });
        setTasks(tasks.map(task => (task.id === id ? updatedTask : task)));
      } catch (error) {
        console.error("Error updating task:", error);
      }
    }
  };

  const handleDeleteTask = async (id) => {
    try {
      await deleteTask(id);
      setTasks(tasks.filter(task => task.id !== id));
    } catch (error) {
      console.error("Error deleting task:", error);
    }
  };

  return (
    <div className="App">
      <h1>Task Manager</h1>
      <AddTaskForm onAddTask={handleAddTask} />
      <TaskList
        tasks={tasks}
        onToggleComplete={handleToggleComplete}
        onDelete={handleDeleteTask}
      />
    </div>
  );
}

export default App;