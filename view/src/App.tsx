import React, {useEffect, useState} from 'react';
import HomePage from './HomePage';
import LoginPage from './LoginPage';

interface Props {}

const App = (props: Props) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  useEffect(() => {
    fetch('http://localhost:8081/api/checklogin', {
      //credentials: 'include',
      mode: 'no-cors',
    })
      .then(response => response.json())
      .then(data => {
        const {is_logged_in} = data;
        setIsLoggedIn(is_logged_in);
      })
      .catch(error => console.log(error));
  }, []);
  return (
    //isLoggedIn ? <HomePage /> : <LoginPage />
    <HomePage />
  );
};

export default App;
