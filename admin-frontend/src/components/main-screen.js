import React from 'react';
import { Button, Tab } from 'semantic-ui-react';
import styled from 'styled-components/macro';
import AnnouncementForm from './announcement-form';

const LogoutButtonContainer = styled.div`
  display: flex;
  justify-content: flex-end;
  margin-bottom: 8px;
`;

const tabPanes = [
  {
    menuItem: 'Announcements',
    render: () => (
      <Tab.Pane>
        <AnnouncementForm />
      </Tab.Pane>
    )
  },
  {
    menuItem: 'Tweets',
    render: () => <Tab.Pane>Tweets</Tab.Pane>
  }
];

const MainScreen = ({ history }) => (
  <>
    <LogoutButtonContainer>
      <Button
        negative
        onClick={() => {
          localStorage.removeItem('token');
          history.push('/login');
        }}
      >
        Logout
      </Button>
    </LogoutButtonContainer>
    <Tab panes={tabPanes} />
  </>
);

export default MainScreen;
