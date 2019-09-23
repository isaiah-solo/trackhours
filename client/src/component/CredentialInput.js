import React, {useState} from 'react';
import styled from 'styled-components'

function CredentialInput({onChange, placeholder, type, value}) (
  <Input
    onChange={onChange}
    placeholder={placeholder}
    type={type}
    value={value} />
);

const Input = styled.input`
  border-radius: 4px;
  border-width: 0px;
  box-sizing: border-box;
  font-size: 24px;
  margin-top: 16px;
  padding: 8px;
  width: 300px;
  :focus {
    outline-width: 0;
  }
`;

export default CredentialInput;
