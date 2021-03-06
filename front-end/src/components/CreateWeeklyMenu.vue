<template>
  <v-dialog
    :value="dialog"
    persistent
    :disabled="disabledDialog"
    max-width="500px"
  >
    <v-card color="rgb(210, 248, 149)">
      <v-card-title>
        <span class="text-h5">{{ isEditMode ? "Edit Menu" : "New Menu" }}</span>
      </v-card-title>

      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12" sm="6" md="6">
              <v-autocomplete
                v-model="weeklyMenu.numPeople"
                :items="[2, 4]"
                filled
                label="Number of people"
              ></v-autocomplete>
            </v-col>
            <v-col cols="12" sm="6" md="6">
              <v-autocomplete
                v-model="weeklyMenu.recipesPerWeek"
                :items="[3, 4, 5]"
                filled
                label="Recipes per week"
              ></v-autocomplete>
            </v-col>
            <v-row justify="space-around">
              <v-col cols="3" sm="3" md="3">
                <label for="">Starting week</label>
              </v-col>
              <v-col cols="12" sm="12" md="12">
                <v-row justify="space-around">
                  <v-date-picker
                    header-color="rgb(6, 122, 70)"
                    :allowed-dates="allowedDates"
                    v-model="weeklyMenu.firstDayWeek"
                  ></v-date-picker>
                </v-row>
              </v-col>
            </v-row>
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
          :disabled="!isValid"
        >
          Create
        </v-btn>
        <v-btn
          v-else
          color="blue darken-1"
          text
          @click="edit"
          :disabled="!isValid"
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
import moment from "moment";

const weeklyMenuBase = {
      numPeople: 2,
      recipesPerWeek: 3,
      firstDayWeek: undefined,
    }

export default Vue.extend({
  name: "CreateWeeklyMenu",
  props: {
    dialog: Boolean,
    itemToUpdate: Object,
  },
  watch: {
    itemToUpdate(val) {
      if (val !== undefined) {
        this.weeklyMenu = {
          numPeople: val.numPeople,
          recipesPerWeek: val.daysPerWeek,
          firstDayWeek: val.firstDayWeek,
        };
      }
    },
  },
  methods: {
    close() {
      this.$emit("close-dialog");
    },
    async edit() {
      try {
        this.disabledDialog = true;
        await (this as any).$http.put(`/weeklyMenus/${this.itemToUpdate.id}`, {
          numPeople: this.weeklyMenu.numPeople,
          daysPerWeek: this.weeklyMenu.recipesPerWeek,
          firstDayWeek: this.weeklyMenu.firstDayWeek,
        });
        this.$emit("item-updated");
        this.weeklyMenu = {...weeklyMenuBase}
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
        await (this as any).$http.post(`/weeklyMenus`, {
          numPeople: this.weeklyMenu.numPeople,
          daysPerWeek: this.weeklyMenu.recipesPerWeek,
          firstDayWeek: this.weeklyMenu.firstDayWeek,
        });
        this.$emit("item-created");
        this.weeklyMenu = {...weeklyMenuBase}
        this.close();
      } catch (error) {
        this.apiError = error;
        this.snackbar = true;
      } finally {
        this.disabledDialog = false;
      }
    },
    allowedDates: (val: string) => moment(val).day() === 1,
  },
  computed: {
    isValid() {
      return (
        this.weeklyMenu.numPeople !== undefined &&
        this.weeklyMenu.recipesPerWeek !== undefined &&
        this.weeklyMenu.firstDayWeek !== undefined
      );
    },
    isEditMode() {
      return this.itemToUpdate !== undefined;
    },
  },
  data: () => ({
    apiError: "",
    snackbar: false,
    disabledDialog: false,
    weeklyMenu: {
      ...weeklyMenuBase
    }
  }),
});
</script>
