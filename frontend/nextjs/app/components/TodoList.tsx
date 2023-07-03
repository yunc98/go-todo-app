import { Todo } from '../types/todo';
import TodoField from './TodoField';

export default async function TodoList({
  promise,
}: {
  promise: Promise<Todo[]>;
}) {
  const todos = await promise;

  return (
    <div>
      {todos.map((todo) => (
        <TodoField todo={todo} />
      ))}
    </div>
  );
}
