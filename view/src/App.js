import React, {useEffect, useState} from 'react';
import HomePage from './HomePage';
import LoginPage from './LoginPage';

const App = (props) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  useEffect(() => {
    fetch('http://trackhours.co/api/checklogin', {
      credentials: 'include',
    })
      .then(response => response.json())
      .then(data => {
        const {is_logged_in} = data;
        setIsLoggedIn(is_logged_in);
      })
      .catch(error => console.log(error));
  }, []);
  return (
    isLoggedIn ? <HomePage /> : <LoginPage />
  );
};

export default App;
