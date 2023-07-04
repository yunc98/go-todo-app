'use client';

import { useRouter } from 'next/navigation';
import { Todo } from '../types/todo';

export default function TodoField({ todo }: { todo: Todo }) {
  const router = useRouter();

  const getNewStatus = () => {
    if (todo.completed) {
      return 'false';
    } else {
      return 'true';
    }
  };

  const handleCompleted = async () => {
    const data = new URLSearchParams({
      title: todo.title,
      completed: getNewStatus(),
    });

    await fetch(`http://localhost:8080/todo/${todo.ID}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: data,
    })
      .then(() => {
        router.refresh();
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const handleDelete = async () => {
    await fetch(`http://localhost:8080/todo/${todo.ID}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
    })
      .then(() => {
        router.refresh();
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <div className="flex justify-between">
      <div className="flex gap-8">
        <button onClick={() => handleCompleted()}>
          {todo.completed ? '✅' : '⬜'}
        </button>
        <p>{todo.title}</p>
      </div>
      <button onClick={() => handleDelete()} className="text-xs">
        ❌
      </button>
    </div>
  );
}
