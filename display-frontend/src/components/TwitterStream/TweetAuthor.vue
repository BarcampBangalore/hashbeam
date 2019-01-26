<template>
  <div :class="$style.grid">
    <img :class="$style.picture" :src="highResImage">
    <h1 :class="$style.name">{{ user.name }}</h1>
    <h2 :class="$style.screenName">@{{ user.screen_name }}</h2>
  </div>
</template>

<script>
export default {
  props: ['user'],

  computed: {
    highResImage() {
      const url = this.user.profile_image_url_https;
      if (url.endsWith('jpeg')) {
        return `${url.slice(0, -12)}${url.slice(-5)}`;
      }

      return `${url.slice(0, -11)}${url.slice(-4)}`;
    }
  }
};
</script>

<style module>
.grid {
  display: grid;
  grid-template-columns: 3fr 11fr;
  grid-template-rows: 2fr 2fr;
  grid-template-areas:
    'picture name'
    'picture screenName';
}

.picture {
  grid-area: picture;
  justify-self: end;
  align-self: center;
  object-fit: cover;
  height: calc((11.5vh + 11.5vw) / 2);
  width: calc((11.5vh + 11.5vw) / 2);
  border-radius: 100%;
  margin-right: 10%;
}

.name {
  grid-area: name;
  color: white;
  font-size: calc((3.5vh + 3.5vw) / 2);
  font-weight: bold;
  align-self: end;
}
.screenName {
  grid-area: screenName;
  color: rgba(255, 255, 255, 0.6);
  margin-top: 1%;
  font-size: calc((3vh + 2vw) / 2);
}
</style>
