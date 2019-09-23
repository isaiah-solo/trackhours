import React, {useState} from 'react';
import styled from 'styled-components'

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

interface Props {
  onChange: (e: any) => void,
  placeholder: string,
  type: string,
  value: string,
}

const CredentialInput = ({onChange, placeholder, type, value}: Props) => (
  <Input
    onChange={onChange}
    placeholder={placeholder}
    type={type}
    value={value} />
);

export default CredentialInput;
