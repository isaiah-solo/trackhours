import React from 'react';
import styled from 'styled-components';

import Page from './component/Page';

interface Props {}

function HomePage(props: Props) {
  return (
    <Page>
      <Sidebar>
        <SidebarItem />
        <SidebarItem />
      </Sidebar>
      <Content>
        Content
      </Content>
    </Page>
  );
};

const Content = styled.div`
  align-items: center;
  background-color: #202020;
  box-sizing: border-box;
  color: white;
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  height: 100%;
`;

const Sidebar = styled.div`
  background-color: #202020;
  box-sizing: border-box;
  color: white;
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 20px;
  width: 300px;
`;

const SidebarItem = styled.div`
  background-color: #ee0060;
  box-sizing: border-box;
  color: white;
  height: 50px;
  margin-top: 20px;
  width: 100%;
  &:first-child {
    margin-top: 0;
  }
`;

export default HomePage;
