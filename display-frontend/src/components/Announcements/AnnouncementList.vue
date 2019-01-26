<template>
  <div :class="[$style.grid, $style.background]">
    <h1 :class="$style.title">Announcements</h1>
    <announcement-item
      v-for="announcement in announcements.slice(0, 8)"
      :key="announcement.id"
      :announcement="announcement"
    ></announcement-item>
  </div>
</template>

<script>
import AnnouncementItem from "./AnnouncementItem";
import gql from "graphql-tag";

export default {
  components: {
    AnnouncementItem
  },

  data() {
    return {
      announcements: []
    };
  },

  async created() {
    const response = await this.$apollo.provider.defaultClient.query({
      query: gql`
        query {
          announcements {
            id
            timestampISO8601
            message
          }
        }
      `,
      fetchPolicy: "network-only"
    });

    this.announcements = response.data.announcements;
  },

  sockets: {
    announcement(announcement) {
      this.announcements.unshift(announcement);
    }
  }
};
</script>

<style module>
.grid {
  display: grid;
  grid-template-rows: 10vh;
  grid-auto-rows: max-content;
  grid-row-gap: 5vh;
}

.background {
  background-color: rgba(0, 0, 0, 0.5);
}

.title {
  color: white;
  font-size: calc((3.2vh + 3.2vw) / 2);
  font-weight: bold;
  align-self: center;
  margin-left: 0.5em;
}
</style>
