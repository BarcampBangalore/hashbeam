<template>
  <div :class="$style.announcementItem">
    <p :class="$style.timestring">{{ timestring }}</p>
    <p :class="$style.message">{{ announcement.message }}</p>
  </div>
</template>

<script>
export default {
  props: ['announcement'],

  computed: {
    timestring() {
      const date = new Date(this.announcement.datetime);

      return date.toLocaleTimeString('en-US', {
        hour12: true,
        hour: '2-digit',
        minute: '2-digit'
      });
    }
  }
};
</script>

<style module>
.timestring {
  font-size: calc((1.5vh + 1.5vw) / 2);
  color: rgba(255, 255, 255, 0.6);
  margin: 0 0 1vh 1em;
}

.message {
  font-size: calc((1.8vh + 1.8vw) / 2);
  color: white;
  margin: 0 0.5em 0em 0.8em;
  line-height: 1.2em;
}
</style>
