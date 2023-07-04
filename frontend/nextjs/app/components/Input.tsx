'use client';

import { useRouter } from 'next/navigation';
import { useState } from 'react';

export default function Input() {
  const router = useRouter();
  const [newTodo, setNewTodo] = useState('');

  const onChangeNewTodo = (
    event: React.ChangeEvent<HTMLInputElement>
  ): void => {
    setNewTodo(event.target.value);
  };

  const onClickAddNewTodo = async () => {
    const data = new URLSearchParams({
      title: newTodo,
      completed: 'false',
    });

    await fetch('http://localhost:8080/todo', {
      method: 'POST',
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

    setNewTodo('');
  };

  return (
    <div>
      <input
        type="text"
        placeholder="Add a new todo"
        value={newTodo}
        onChange={onChangeNewTodo}
      />
      <button onClick={onClickAddNewTodo}>Add</button>
    </div>
  );
}
