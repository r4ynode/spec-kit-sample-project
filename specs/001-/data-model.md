# Data Model

**Feature**: Simple Task Management Application

## Entities

### Task
Represents a single to-do item.

**Schema**:

| Field       | Type    | Description                               |
|-------------|---------|-------------------------------------------|
| `id`        | Integer | **Primary Key**. Unique identifier for the task. |
| `title`     | String  | The description or content of the task.   |
| `completed` | Boolean | Indicates if the task is complete. `false` by default. |
| `created_at`| Timestamp| The date and time when the task was created. |
