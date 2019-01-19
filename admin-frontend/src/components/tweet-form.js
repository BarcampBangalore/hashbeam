import gql from 'graphql-tag';
import React from 'react';
import { Mutation, Query } from 'react-apollo';
import Tweet from 'react-tweet';
import { Button, Icon, Loader, Message } from 'semantic-ui-react';
import styled from 'styled-components/macro';
import randomEmoji from '../util/random-emoji';

const FETCH_TWEET_TO_REVIEW_QUERY = gql`
  query FetchTweetToReview {
    tweetToReview {
      tweetId
      authorId
      timestampISO8601
      originalTweetJson
    }
  }
`;

const WHITELIST_TWEETER_MUTATION = gql`
  mutation WhitelistTweeter($twitterAuthorID: String!) {
    whitelistTweeter(twitterAuthorID: $twitterAuthorID)
  }
`;

const MUTE_TWEETER_MUTATION = gql`
  mutation MuteTweeter($twitterAuthorID: String!) {
    muteTweeter(twitterAuthorID: $twitterAuthorID)
  }
`;

const ButtonsContainer = styled.div`
  display: flex;
  width: 100%;
  margin-bottom: 8px;
`;

const ReloadButton = styled(Button).attrs({
  primary: true
})`
  &&& {
    width: 25px;
    display: flex;
    justify-content: center;
  }
`;

const ReloadIcon = styled(Icon).attrs({
  name: 'refresh'
})`
  margin: 0px !important;
`;

const ActionButton = styled(Button)`
  flex: 1;
`;

const TweetContainer = styled.div`
  width: 100%;
  min-height: 64px;
  display: flex;
  justify-content: center;
  align-items: center;
`;

const NoTweetsToReviewContainer = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
`;

const Emoji = styled.div`
  font-size: 64px;
  margin-top: 16px;
  margin-bottom: 16px;
`;

const Loading = () => <Loader active inline="centered" />;

const ErrorMessage = styled(Message).attrs({
  negative: true
})`
  width: 100%;
  display: flex;
  justify-content: center;
`;

const ApproveButton = props => <ActionButton positive {...props} />;
const RejectButton = props => <ActionButton negative {...props} />;

const TweetForm = () => (
  <Query query={FETCH_TWEET_TO_REVIEW_QUERY} notifyOnNetworkStatusChange={true}>
    {({ loading: tweetLoading, error, data, refetch }) => (
      <Mutation
        mutation={WHITELIST_TWEETER_MUTATION}
        onCompleted={() => refetch()}
      >
        {(whitelistTweeter, { loading: whitelistTweeterLoading }) => (
          <Mutation
            mutation={MUTE_TWEETER_MUTATION}
            onCompleted={() => refetch()}
          >
            {(muteTweeter, { loading: muteTweeterLoading }) => (
              <>
                <ButtonsContainer>
                  <ReloadButton
                    onClick={() => refetch()}
                    disabled={tweetLoading}
                  >
                    <ReloadIcon />
                  </ReloadButton>
                  <RejectButton
                    onClick={() =>
                      muteTweeter({
                        variables: {
                          twitterAuthorID: data.tweetToReview.authorId
                        }
                      })
                    }
                    loading={muteTweeterLoading}
                    disabled={Boolean(
                      tweetLoading ||
                        whitelistTweeterLoading ||
                        muteTweeterLoading ||
                        !data.tweetToReview ||
                        error
                    )}
                  >
                    {muteTweeterLoading || <Icon name="close" />}
                  </RejectButton>
                  <ApproveButton
                    onClick={() =>
                      whitelistTweeter({
                        variables: {
                          twitterAuthorID: data.tweetToReview.authorId
                        }
                      })
                    }
                    loading={whitelistTweeterLoading}
                    disabled={Boolean(
                      tweetLoading ||
                        whitelistTweeterLoading ||
                        muteTweeterLoading ||
                        !data.tweetToReview ||
                        error
                    )}
                  >
                    {whitelistTweeterLoading || <Icon name="check" />}
                  </ApproveButton>
                </ButtonsContainer>
                <TweetContainer>
                  {(() => {
                    if (tweetLoading) {
                      return <Loading />;
                    }

                    if (error) {
                      console.error(error);
                      return (
                        <ErrorMessage>Error logged to console</ErrorMessage>
                      );
                    }

                    if (!data.tweetToReview) {
                      const { name, text } = randomEmoji();
                      return (
                        <NoTweetsToReviewContainer>
                          <h2>No tweets to review right now</h2>
                          <p>Here's a {name} instead:</p>
                          <Emoji>{text}</Emoji>
                        </NoTweetsToReviewContainer>
                      );
                    }

                    const tweetJson = JSON.parse(
                      data.tweetToReview.originalTweetJson
                    );

                    return <Tweet data={tweetJson} />;
                  })()}
                </TweetContainer>
              </>
            )}
          </Mutation>
        )}
      </Mutation>
    )}
  </Query>
);

export default TweetForm;
