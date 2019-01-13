import gql from 'graphql-tag';
import React from 'react';
import { Mutation } from 'react-apollo';
import { Button, Form, Icon, TextArea } from 'semantic-ui-react';
import styled from 'styled-components/macro';

const SubmitButton = styled(Button).attrs({
  primary: true
})`
  &&& {
    margin-top: 16px;
    width: 100%;
  }
`;

const MAKE_ANNOUNCEMENT_MUTATION = gql`
  mutation MakeAnnouncement($announcementText: String!) {
    makeAnnouncement(message: $announcementText) {
      id
    }
  }
`;

export default class Announcement extends React.Component {
  state = {
    announcementText: '',
    showSuccess: false,
    showError: false
  };

  render() {
    return (
      <Mutation
        mutation={MAKE_ANNOUNCEMENT_MUTATION}
        onCompleted={() => {
          this.setState({ announcementText: '', showSuccess: true });
          setTimeout(() => this.setState({ showSuccess: false }), 3000);
        }}
        onError={() => {
          this.setState({ announcementText: '', showError: true });
          setTimeout(() => this.setState({ showError: false }), 3000);
        }}
      >
        {(makeAnnouncement, { loading }) => (
          <Form
            onSubmit={e => {
              e.preventDefault();
              if (!this.state.announcementText) {
                return;
              }
              makeAnnouncement({
                variables: { announcementText: this.state.announcementText }
              });
            }}
          >
            <TextArea
              placeholder="Announcement text"
              value={this.state.announcementText}
              onChange={e =>
                this.setState({ announcementText: e.target.value })
              }
            />
            <SubmitButton
              type="submit"
              loading={loading}
              positive={this.state.showSuccess}
              negative={this.state.showError}
            >
              {(this.state.showSuccess && <Icon name="check" />) ||
                (this.state.showError && <Icon name="close" />) ||
                'Submit'}
            </SubmitButton>
          </Form>
        )}
      </Mutation>
    );
  }
}
