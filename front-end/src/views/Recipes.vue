<template>
  <recipes-manager />
</template>

<script lang="ts">
import Vue from "vue";
import RecipesManager from "../components/Recipes.vue";

export default Vue.extend({
  name: "Recipes",

  components: {
    RecipesManager,
  },
});
</script>
