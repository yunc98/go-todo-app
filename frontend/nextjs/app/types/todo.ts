export type Todo = {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  title: string;
  completed: boolean;
};
