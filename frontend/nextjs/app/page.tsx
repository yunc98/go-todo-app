import TodoList from './components/TodoList';

const getAllTodos = async () => {
  try {
    const res = await fetch('http://backend:8080/todo', { cache: 'no-store' });
    return res.json();
  } catch (error) {
    console.log(error);
  }
};

export default async function Home() {
  const todosData = getAllTodos();

  return (
    <main>
      <div className="z-10 w-full max-w-5xl items-center justify-between">
        <h1 className="text-center">Todo list😁</h1>
        <TodoList promise={todosData} />
      </div>
    </main>
  );
}
