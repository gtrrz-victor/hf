<template>
  <v-dialog
    :value="dialog"
    persistent
    :disabled="disabledDialog"
    max-width="500px"
  >
    <v-card color="rgb(210, 248, 149)">
      <v-card-title>
        <span class="text-h5">{{
          isEditMode ? "Edit Recipe" : "New Recipe"
        }}</span>
      </v-card-title>

      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12" sm="6" md="6">
              <v-text-field label="Title" v-model="recipe.title"></v-text-field>
            </v-col>
            <v-col cols="12" sm="6" md="6">
              <v-text-field
                label="Subtitle"
                v-model="recipe.subtitle"
              ></v-text-field>
            </v-col>
            <v-col cols="12" sm="6" md="6">
              <v-text-field
                label="Difficulty"
                v-model="recipe.difficulty"
              ></v-text-field>
            </v-col>
            <v-col cols="12" sm="6" md="6">
              <v-text-field
                label="Cook time"
                v-model="recipe.cook_time"
                :rules="[rules.greatherThanZero]"
                suffix="minutes"
                type="number"
              ></v-text-field>
            </v-col>
            <v-col cols="12" sm="12" md="12">
              <v-text-field label="Image" v-model="recipe.image"></v-text-field>
            </v-col>
            <v-col cols="12" sm="12" md="12">
              <v-textarea
                label="Description"
                v-model="recipe.description"
              ></v-textarea>
            </v-col>
            <v-col cols="12" sm="12" md="12">
              <v-combobox
                multiple
                v-model="recipe.tags"
                label="Tags"
                append-icon
                chips
                deletable-chips
                class="tag-input"
              >
              </v-combobox>
            </v-col>
            <v-col cols="12" sm="12" md="12">
              <v-combobox
                multiple
                v-model="recipe.utensils"
                label="Utensils"
                append-icon
                chips
                deletable-chips
                class="tag-input"
              >
              </v-combobox>
            </v-col>
            <v-col cols="12" sm="12" md="12">
              <v-combobox
                multiple
                v-model="recipe.allergens"
                label="Allergens"
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
        <v-btn color="blue darken-1" text @click="close"> Cancel </v-btn>
        <v-btn
          v-if="!isEditMode"
          color="blue darken-1"
          text
          @click="save"
          :disabled="isEmpty(recipe.title)"
        >
          Create
        </v-btn>
        <v-btn
          v-else
          color="blue darken-1"
          text
          @click="edit"
          :disabled="isEmpty(recipe.title)"
        >
          Edit
        </v-btn>
      </v-card-actions>
    </v-card>
    <v-snackbar v-model="snackbar">
      {{ apiError }}
    </v-snackbar>
  </v-dialog>
</template>


<script lang="ts">
import Vue from "vue";

const recipeBase = {
  difficulty: "",
  title: "",
  description: "",
  image: "",
  subtitle: "",
  cook_time: 1,
  tags: [],
  utensils: [],
  allergens: [],
};

export default Vue.extend({
  name: "CreateRecipe",
  props: {
    dialog: Boolean,
    itemToUpdate: Object,
  },
  watch: {
    itemToUpdate(val) {
      if (val !== undefined) {
        this.recipe = {
          difficulty: val.difficulty,
          title: val.title,
          description: val.description,
          image: val.image,
          subtitle: val.subtitle,
          cook_time: val.cook_time,
          tags: val.tags,
          utensils: val.utensils,
          allergens: val.allergens,
        };
      }
    },
  },
  methods: {
    async edit() {
      try {
        this.disabledDialog = true;
        await (this as any).$http.put(`/recipes/${this.itemToUpdate.id}`, {
          ...this.recipe,
        });
        this.$emit("item-updated");
        this.recipe = { ...recipeBase };
        this.close();
      } catch (error) {
        this.apiError = error;
        this.snackbar = true;
      } finally {
        this.disabledDialog = false;
      }
    },
    async save() {
      try {
        this.disabledDialog = true;
        await (this as any).$http.post(`/recipes`, { ...this.recipe,cook_time : parseInt(this.recipe.cook_time+"") });
        this.$emit("item-created");
        this.recipe = { ...recipeBase };
        this.close();
      } catch (error) {
        this.apiError = error;
        this.snackbar = true;
      } finally {
        this.disabledDialog = false;
      }
    },
    isEmpty: (val: string) => val.length === 0,
    close() {
      this.$emit("close-dialog");
    },
  },
  computed: {
    isEditMode() {
      return this.itemToUpdate !== undefined;
    },
  },
  data: () => ({
    rules: {
      greatherThanZero: (value: number) => value > 0 || "Greather than Zero.",
    },
    apiError: "",
    snackbar: false,
    disabledDialog: false,
    recipe: {
      ...recipeBase,
    },
  }),
});
</script>
