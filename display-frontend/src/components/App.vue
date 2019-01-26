<template>
  <div :class="[$style.grid, $style.background]">
    <twitter-stream></twitter-stream>
    <announcement-list></announcement-list>
  </div>
</template>

<script>
import TwitterStream from './TwitterStream/TwitterStream';
import AnnouncementList from './Announcements/AnnouncementList';

export default {
  components: {
    TwitterStream,
    AnnouncementList
  }
};
</script>

<style>
body {
  margin: 0;
  font-family: 'Fira Sans', sans-serif;
}
</style>

<style module>
.grid {
  display: grid;
  height: 100vh;
  grid-template-columns: 9fr 4fr;
}

.background {
  background: linear-gradient(to right, #f2994a, #f2c94c);
}
</style>
