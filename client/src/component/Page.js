import React from 'react';
import styled from 'styled-components'

function Page({
  children
}) {
  return (
    <Root>
      {children}
    </Root>
  );
};

const Root = styled.div`
  background-color: #202020;
  box-sizing: border-box;
  color: white;
  display: flex;
  flex-direction: row;
  height: 100vh;
  width: 100vw;
`;

export default React.memo(Page);
