<template>
  <weekly-menu />
</template>

<script lang="ts">
import Vue from "vue";
import WeeklyMenu from "../components/WeeklyMenu.vue";

export default Vue.extend({
  name: "Home",

  components: {
    WeeklyMenu,
  },
});
</script>
