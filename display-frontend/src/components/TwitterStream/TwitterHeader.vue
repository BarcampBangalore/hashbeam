<template>
  <div :class="$style.grid">
    <img src="../../assets/logo.png" :class="$style.logo">
    <h1 :class="$style.trackText">{{ trackText }}</h1>
  </div>
</template>

<script>
export default {
  props: ["trackText"]
};
</script>

<style module>
.grid {
  display: grid;
  grid-template-columns: 1fr 3fr;
  grid-template-areas: "logo trackText";
}

.logo {
  grid-area: logo;
  justify-self: center;
  align-self: center;
  height: calc((14vh + 14vw) / 2);
}

.trackText {
  grid-area: trackText;
  color: rgba(255, 255, 255, 0.8);
  font-size: calc((4vh + 4vw) / 2);
  align-self: center;
  justify-self: end;
  margin-right: 1em;
}
</style>
