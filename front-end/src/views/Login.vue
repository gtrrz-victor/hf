<template>
  <v-container>
    <v-row justify="space-around" class="pt-5">
      <v-col cols="3" sm="3" md="3">
        <v-card>
          <v-img
            height="250"
            src="https://img.hellofresh.com/c_fill,f_auto,fl_lossy,h_650,q_auto,w_1199/hellofresh_website/gb/cms/homepage/HF_App_web_landing_page_newbanners_desktop_1199x650.jpg"
          ></v-img>
          <v-card-title> Welcome to Hello Fresh </v-card-title>
          <v-card-subtitle>
            <v-row justify="space-around" class="pt-5">
              <v-col cols="12" sm="12" md="12">
                <div>If you want to manage your weekly plans/recipes.</div>
              </v-col>
            </v-row>
            <v-row justify="space-around" class="pt-5">
              <v-col cols="5" sm="5" md="5">
                <GoogleLogin
                  :params="params"
                  :onSuccess="onSuccess"
                  :onFailure="onFailure"
                  :renderParams="renderParams"
                ></GoogleLogin>
              </v-col>
            </v-row>
          </v-card-subtitle>
          <v-card-actions> </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import GoogleLogin from "vue-google-login";

export default Vue.extend({
  name: "Login",

  components: {
    GoogleLogin,
  },
  methods: {
    onSuccess(data: any) {
      this.$cookies.set("user", JSON.stringify(data));
      this.$router.push("/");
    },
    onFailure(data: any) {
      alert(data);
    },
  },
  data: () => ({
    params: {
      client_id:
        "487249360034-5bq5e0t0a1e2fellutlemiob0f3h7the.apps.googleusercontent.com",
    },
    renderParams: {
      width: 150,
      height: 40,
    },
  }),
});
</script>
