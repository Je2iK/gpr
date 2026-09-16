import { UserForm } from '../features/users/UserForm';

export const MainPage = () => {
  return (
    <main style={{ padding: '40px' }}>
      <h1>Главная страница</h1>
      <UserForm />
    </main>
  );
};