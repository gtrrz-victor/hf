<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-data-table
          :headers="headers"
          :items="recipes"
          :items-per-page="5"
          class="elevation-1"
        >
          <template v-slot:top>
            <v-toolbar flat v-if="!isFilterEnabled">
              <v-toolbar-title>Recipes CRUD</v-toolbar-title>
              <v-divider class="mx-4" inset vertical></v-divider>
              <v-spacer></v-spacer>
              <v-btn class="mb-2" @click="dialog = true"> New Recipe </v-btn>
              <create-recipe
                :dialog="dialog"
                :itemToUpdate="itemToUpdate"
                @close-dialog="close"
                @item-created="getData"
                @item-updated="getData"
              >
              </create-recipe>
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
import CreateRecipe from "./CreateRecipe.vue";

export default Vue.extend({
  name: "Recipes",
  props: {
    menuID: Number,
  },
  components: {
    CreateRecipe,
  },
  computed: {
    isFilterEnabled(): boolean {
      return this.menuID !== undefined;
    },
  },
  watch: {
    menuID() {
      this.getData();
    },
  },
  methods: {
    editItem(item: any) {
      this.itemToUpdate = item;
      this.dialog = true;
    },

    deleteItem(item: any) {
      this.idToRemove = item.id;
      this.dialogDelete = true;
    },

    async deleteItemConfirm() {
      try {
        this.disabledDialogDelete = true;
        await (this as any).$http.delete(`/recipes/${this.idToRemove}`);
        this.idToRemove = undefined;
        this.getData();
      } catch (error) {
        console.log(error);
      } finally {
        this.closeDelete();
        this.disabledDialogDelete = false;
      }
      this.closeDelete();
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
        let response;
        if (this.isFilterEnabled) {
          response = await (this as any).$http.get(
            `weeklyMenus/${this.menuID}/recipes`
          );
        } else {
          response = await (this as any).$http.get("/recipes");
        }
        if (response.status === 204) {
          this.recipes = [];
        } else {
          this.recipes = response.data;
        }
      } catch (error) {
        console.log(error);
      }
    },
  },
  created() {
    this.getData();
  },
  data: () => ({
    itemToUpdate: undefined,
    idToRemove: undefined,
    disabledDialogDelete: false,
    dialogDelete: false,
    dialog: false,
    headers: [
      {
        text: "Recipe Id",
        align: "start",
        value: "id",
      },
      { text: "Difficulty", value: "difficulty" },
      { text: "Description", value: "description" },
      { text: "Image", value: "image" },
      { text: "Title", value: "title" },
      { text: "Subtitle", value: "subtitle" },
      { text: "Cook time", value: "cook_time" },
      { text: "Tags", value: "tags" },
      { text: "Utensils", value: "utensils" },
      { text: "Allergens", value: "allergens" },
      { text: "Actions", value: "actions", sortable: false },
    ],
    recipes: [],
  }),
});
</script>
