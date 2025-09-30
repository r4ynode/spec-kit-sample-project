import React from 'react';

const TaskItem = ({ task, onToggleComplete, onDelete }) => {
  return (
    <div classNameName={`task-item ${task.completed ? 'completed' : ''}`}>
      <input
        type="checkbox"
        checked={task.completed}
        onChange={() => onToggleComplete(task.id)}
      />
      <span>{task.title}</span>
      <button onClick={() => onDelete(task.id)}>Delete</button>
    </div>
  );
};

export default TaskItem;
