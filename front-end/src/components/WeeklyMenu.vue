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
              <v-btn class="mb-2" @click="dialog = true"> New Menu </v-btn>
              <create-weekly-menu
                :itemToUpdate="itemToUpdate"
                :dialog="dialog"
                @close-dialog="close"
                @item-created="getData"
                @item-updated="getData"
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

              <v-dialog
                v-model="dialogAddRecipes"
                persistent
                :disabled="disabledDialogAddRecipes"
                max-width="500px"
              >
                <v-card>
                  <v-card-title class="text-h5"
                    >Add/Remove recipes</v-card-title
                  >
                  <v-card-text>
                    <v-container>
                      <v-row>
                        <v-col cols="12" sm="6" md="6">
                          <v-combobox
                            multiple
                            v-model="selectedWeek.recipesMenu"
                            label="Recipe Ids"
                            append-icon
                            chips
                            deletable-chips
                            class="tag-input"
                          >
                          </v-combobox>
                        </v-col>
                      </v-row>
                    </v-container>
                  </v-card-text>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="closeAddRecipes"
                      >Cancel</v-btn
                    >
                    <v-btn color="blue darken-1" text @click="updateRecipes"
                      >Edit</v-btn
                    >
                    <v-spacer></v-spacer>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </v-toolbar>
          </template>
          <template v-slot:[`item.actions`]="{ item }">
            <v-icon small class="mr-2" @click="loadMenuRecipes(item)">
              mdi-plus-box-multiple
            </v-icon>
            <v-icon small class="mr-2" @click="editItem(item)">
              mdi-pencil
            </v-icon>
            <v-icon small @click="deleteItem(item)"> mdi-delete </v-icon>
          </template></v-data-table
        >
      </v-col>
    </v-row>
    <v-snackbar v-model="snackbar">
      {{ apiError }}
    </v-snackbar>
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
    closeAddRecipes() {
      this.selectedWeek = {
        recipesMenu: [],
        id: "",
      };
      this.dialogAddRecipes = false;
    },
    async loadMenuRecipes(item: any) {
      try {
        const { data } = await (this as any).$http.get(
          `/weeklyMenus/${item.id}/recipes`
        );
        this.selectedWeek = {
          recipesMenu: data.map((recipe: any) => recipe.id),
          id: item.id,
        };
        this.dialogAddRecipes = true;
      } catch (err) {
        this.apiError = err;
        this.snackbar = true;
      }
    },
    async updateRecipes() {
      try {
        this.disabledDialogAddRecipes = true;
        const { data } = await (this as any).$http.put(
          `/weeklyMenus/${this.selectedWeek.id}/recipes`,
          this.selectedWeek.recipesMenu
        );
        this.closeAddRecipes();
      } catch (err) {
        this.apiError = err;
        this.snackbar = true;
      } finally {
        this.disabledDialogAddRecipes = false;
      }
    },
    editItem(item: any) {
      this.dialog = true;
      this.itemToUpdate = item;
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
      this.itemToUpdate = undefined;
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
    snackbar: false,
    apiError: "",
    dialogAddRecipes: false,
    disabledDialogAddRecipes: false,
    itemToUpdate: undefined,
    idToRemove: undefined,
    disabledDialogDelete: false,
    dialogDelete: false,
    dialog: false,
    selectedWeek: {
      recipesMenu: [],
      id: "",
    },
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
