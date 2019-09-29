import React, {useCallback, useState} from 'react';
import styled from 'styled-components'

import CredentialInput from './component/CredentialInput';
import Page from './component/Page';

import {fetchCreateAccount, fetchLogin} from './api';

function LoginPage(props) {
  const [password, setPassword] = useState('');
  const [username, setUsername] = useState('');
  const createAccount = useCallback(
    () => {
      fetchCreateAccount(
        password,
        username
      ).then((response) => (
        console.log(response.body)
      ));
    },
    [password, username],
  );
  const login = useCallback(
    () => {
      fetchLogin(
        password,
        username
      ).then((response) => (
        console.log(response.body)
      ));
    },
    [password, username],
  );
  return (
    <Page>
      <Content>
        <Title>trackhours</Title>
        <CredentialInput
          onChange={setUsername}
          placeholder="username"
          value={username} />
        <CredentialInput
          onChange={setPassword}
          placeholder="password"
          password
          value={password} />
        <Button onClick={login}>
          Login
        </Button>
        <Button onClick={createAccount}>
          Create Account
        </Button>
      </Content>
    </Page>
  );
};

const Button = styled.div`
  background-color: #ee0060;
  border-radius: 4px;
  box-sizing: border-box;
  color: white;
  cursor: pointer;
  font-size: 24px;
  font-weight: bold;
  margin-top: 16px;
  padding: 8px 0;
  text-align: center;
  user-select: none;
  width: 300px;
`;

const Content = styled.div`
  align-items: center;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  width: 100%;
`;

const Title = styled.div`
  box-sizing: border-box;
  color: white;
  font-size: 64px;
  font-weight: bold;
  text-align: center;
  width: 100%;
`;

export default LoginPage;
