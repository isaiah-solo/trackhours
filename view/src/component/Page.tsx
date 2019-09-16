import React from 'react';
import styled from 'styled-components'

interface Props {
  children: any
}

function Page({children}: Props) {
  return (
    <PageDiv>
      {children}
    </PageDiv>
  );
};

const PageDiv = styled.div`
  background-color: #202020;
  box-sizing: border-box;
  color: white;
  display: flex;
  flex-direction: row;
  height: 100vh;
  width: 100vw;
`;

export default Page;
