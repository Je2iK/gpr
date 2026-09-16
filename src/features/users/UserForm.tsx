import { useState } from 'react';
import { userApi } from './userApi';

export const UserForm = () => {
  const [name, setName] = useState('');
  const [password, setPassword] = useState('');
  const [msg, setMsg] = useState('');

  const handleCreate = async () => {
    try {
      const data = await userApi.create({ name, password });
      setMsg(`Пользователь создан с ID: ${data.id}`);
    } catch (err: any) {
      if (err.response?.status === 409) {
        setMsg('Ошибка: Никнейм уже занят (409)');
      } else {
        setMsg('Ошибка сервера');
      }
    }
  };

  const handleDelete = async () => {
    try {
      await userApi.delete({ name, password });
      setMsg('Пользователь успешно удален');
    } catch (err: any) {
      setMsg('Ошибка при удалении');
    }
  };

  return (
    <div style={{ padding: '20px', border: '1px solid #ccc', borderRadius: '8px', maxWidth: '300px' }}>
      <h3>Управление пользователями</h3>
      <input 
        placeholder="Имя" 
        value={name} 
        onChange={(e) => setName(e.target.value)} 
        style={{ display: 'block', marginBottom: '10px', width: '100%' }}
      />
      <input 
        type="password" 
        placeholder="Пароль" 
        value={password} 
        onChange={(e) => setPassword(e.target.value)} 
        style={{ display: 'block', marginBottom: '10px', width: '100%' }}
      />
      <button onClick={handleCreate} style={{ marginRight: '10px' }}>Создать</button>
      <button onClick={handleDelete}>Удалить</button>
      {msg && <p style={{ marginTop: '10px' }}>{msg}</p>}
    </div>
  );
};