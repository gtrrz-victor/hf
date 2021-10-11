<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-data-table
          :headers="headers"
          :items="weeklyMenu"
          :items-per-page="5"
          class="elevation-1"
        >
          <template v-slot:top>
            <v-toolbar flat>
              <v-toolbar-title>Weekly Menu CRUD</v-toolbar-title>
              <v-divider class="mx-4" inset vertical></v-divider>
              <v-spacer></v-spacer>
              <v-btn class="mb-2" @click="dialog=true">
                New Menu
              </v-btn>
              <create-weekly-menu
                :dialog="dialog"
                @close-dialog="close"
                @item-created="getData"
              >
              </create-weekly-menu>
              <v-dialog
                v-model="dialogDelete"
                persistent
                :disabled="disabledDialogDelete"
                max-width="500px"
              >
                <v-card>
                  <v-card-title class="text-h5"
                    >Are you sure you want to delete this item?</v-card-title
                  >
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="closeDelete"
                      >Cancel</v-btn
                    >
                    <v-btn color="blue darken-1" text @click="deleteItemConfirm"
                      >OK</v-btn
                    >
                    <v-spacer></v-spacer>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </v-toolbar>
          </template>
          <template v-slot:[`item.actions`]="{ item }">
            <v-icon small class="mr-2" @click="editItem(item)">
              mdi-pencil
            </v-icon>
            <v-icon small @click="deleteItem(item)"> mdi-delete </v-icon>
          </template></v-data-table
        >
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import CreateWeeklyMenu from "./CreateWeeklyMenu.vue";

export default Vue.extend({
  name: "WeeklyMenu",
  components: {
    CreateWeeklyMenu,
  },
  methods: {
    editItem(item: any) {
      this.dialog = true;
    },

    deleteItem(item: any) {
      this.idToRemove = item.id;
      this.dialogDelete = true;
    },

    async deleteItemConfirm() {
      try {
        this.disabledDialogDelete = true;
        await (this as any).$http.delete(`/weeklyMenus/${this.idToRemove}`);
        this.idToRemove = undefined;
        this.getData();
      } catch (error) {
        console.log(error);
      } finally {
        this.closeDelete();
        this.disabledDialogDelete = false;
      }
    },

    close() {
      this.dialog = false;
    },

    closeDelete() {
      this.dialogDelete = false;
    },

    async getData() {
      try {
        const response = await (this as any).$http.get("/weeklyMenus");
        this.weeklyMenu = response.data;
      } catch (error) {
        console.log(error);
      }
    },
  },
  created() {
    this.getData();
  },
  data: () => ({
    idToRemove: undefined,
    disabledDialogDelete: false,
    dialogDelete: false,
    dialog: false,
    headers: [
      {
        text: "Menu Id",
        align: "start",
        value: "id",
      },
      { text: "Number of people", value: "numPeople" },
      { text: "Days per week", value: "daysPerWeek" },
      { text: "First day of week", value: "firstDayWeek" },
      { text: "Actions", value: "actions", sortable: false },
    ],
    weeklyMenu: [],
  }),
});
</script>
