<template>
  <v-app>
    <v-app-bar app color="primary">
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Name"
          class="shrink mt-1 hidden-sm-and-down"
          contain
          min-width="100"
          src="https://img.hellofresh.com/image/upload/c_scale,f_auto,q_100,w_435/v1600959926/Hello_Fresh_Lockup.png"
          width="100"
        />
      </div>
      <v-spacer></v-spacer>
      <v-toolbar-title v-if="!!userName()">Welcome {{ userName() }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn to="/recipes" text>
        <span class="mr-2">Recipes</span>
      </v-btn>
      <v-btn to="/" text>
        <span class="mr-2">Weekly Menu</span>
      </v-btn>
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "App",
  methods: {
    userName() {
      const cookie = this.$cookies.get("user");
      if (!cookie) return undefined;
      return (!cookie.it?cookie.dt:cookie.it).Se;
    },
  },
});
</script>
