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
              <v-dialog v-model="dialog" max-width="500px">
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    class="mb-2"
                    v-bind="attrs"
                    v-on="on"
                  >
                    New Menu
                  </v-btn>
                </template>
                <v-card>
                  <v-card-title>
                    <span class="text-h5">New Menu</span>
                  </v-card-title>

                  <v-card-text>
                    <v-container>
                      <v-row>
                        <!-- 
                        <v-col cols="12" sm="6" md="4">
                          <v-text-field
                            v-model="editedItem.name"
                            label="Dessert name"
                          ></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6" md="4">
                          <v-text-field
                            v-model="editedItem.calories"
                            label="Calories"
                          ></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6" md="4">
                          <v-text-field
                            v-model="editedItem.fat"
                            label="Fat (g)"
                          ></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6" md="4">
                          <v-text-field
                            v-model="editedItem.carbs"
                            label="Carbs (g)"
                          ></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6" md="4">
                          <v-text-field
                            v-model="editedItem.protein"
                            label="Protein (g)"
                          ></v-text-field>
                        </v-col>
                        -->
                      </v-row>
                    </v-container>
                  </v-card-text>

                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="close">
                      Cancel
                    </v-btn>
                    <v-btn color="blue darken-1" text @click="save">
                      Save
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
              <v-dialog v-model="dialogDelete" max-width="500px">
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

export default Vue.extend({
  name: "WeeklyMenu",

  methods: {
    editItem(item :any) {
      this.dialog = true;
    },

    deleteItem(item :any) {
      this.dialogDelete = true;
    },

    deleteItemConfirm() {
      this.closeDelete();
    },

    close() {
      this.dialog = false;
    },

    closeDelete() {
      this.dialogDelete = false;
    },

    save() {
      this.close();
    },
  },
  data: () => ({
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
    weeklyMenu: [
    {
        "numPeople": 2,
        "daysPerWeek": 4,
        "firstDayWeek": "2021-10-18",
        "id": 2
    },
    {
        "numPeople": 2,
        "daysPerWeek": 3,
        "firstDayWeek": "2021-10-11",
        "id": 3
    },
    {
        "numPeople": 2,
        "daysPerWeek": 3,
        "firstDayWeek": "2021-10-11",
        "id": 4
    }
],
  }),
});
</script>
