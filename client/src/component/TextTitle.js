import React from 'react';
import styled from 'styled-components'

function TextTitle({children, color = 'white'}) {
  return (
    <Root color={color}>
      {children}
    </Root>
  );
};

const Root = styled.div`
  box-sizing: border-box;
  color: ${({color}) => color};
  font-size: 64px;
  font-weight: bold;
  text-align: center;
  width: 100%;
`;

export default React.memo(TextTitle);
