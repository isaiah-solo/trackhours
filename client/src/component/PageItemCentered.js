import React from 'react';
import styled from 'styled-components'

function PageItemCentered({children}) {
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
  height: fit-content;
  position: relative;
  left: 50%;
  top: 50%;
  transform: perspective(1px) translateY(-50%) translateX(-50%);
`;

export default PageItemCentered;
