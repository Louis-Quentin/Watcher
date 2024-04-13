import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import styles from './css/InscriptionPage.module.css';

const InscriptionPage: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [accepteConditions, setAccepteConditions] = useState(false);
  const [accepteEmails, setAccepteEmails] = useState(false);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log('Email:', email);
    console.log('Password:', password);
    console.log('Accepte conditions:', accepteConditions);
    console.log('Accepte emails:', accepteEmails);
  };

  return (
    <div className={styles.container}>
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
        </div>
        <div className={styles.checkboxGroup}>
          <input
            type="checkbox"
            id="accepteConditions"
            checked={accepteConditions}
            onChange={(e) => setAccepteConditions(e.target.checked)}
          />
          <label htmlFor="accepteConditions">I accept the terms and condition of use</label>
        </div>
        <div className={styles.checkboxGroup}>
          <input
            type="checkbox"
            id="accepteEmails"
            checked={accepteEmails}
            onChange={(e) => setAccepteEmails(e.target.checked)}
          />
          <label htmlFor="accepteEmails">I would like to receive the latest updates       and be notified about nearby watch releases and exclusive offers via newsletter.</label>
        </div>
        <button type="submit" className={styles.submitButton}>Sign up now</button>
        <Link to="/login" className={styles.loginLink}>Log in</Link>
      </form>
    </div>
  );
}

export default InscriptionPage;