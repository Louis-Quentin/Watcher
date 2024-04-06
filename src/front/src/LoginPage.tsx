import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import styles from './css/LoginPage.module.css';

const LoginPage: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log('Email:', email);
    console.log('Password:', password);
    // Ajoutez ici la logique pour la connexion (à implémenter)
  };

  return (
    <div className={styles.container}>
      <Link to="/" className={styles.backButton}>Return</Link>
      <form className={styles.form} onSubmit={handleSubmit}>
        <div className={styles.inputGroup}>
          <label htmlFor="email">Email :</label>
          <input
            type="text"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div className={styles.inputGroup}>
          <label htmlFor="Password">Password :</label>
          <input
            type="password"
            id="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Link to="/inscription" className={styles.forgotLink}>Forgot password</Link>
        </div>
        
        <button type="submit" className={styles.submitButton}>Log in</button>
        <button type="submit" className={styles.appleButton}>Continue with Apple</button>
        <button type="submit" className={styles.googleButton}>Continue with Google</button>
        <Link to="/inscription" className={styles.signupLink}>Sign up now</Link>
      </form>
      <div className={styles.background}></div>
    </div>
  );
}

export default LoginPage;