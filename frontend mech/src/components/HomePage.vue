<template>
  <div class="text-center ma-4">
    <v-row no-gutters>
      <v-col cols="8">
        <v-badge
          color="#C0294A"
          :content="orders"
          :value="orders"
          class="text-h4 pa-4"
          floating
        >
          Incoming order
        </v-badge>
        <v-card @click="dialog = true" variant="tonal" class="text-left pa-4 ma-4">
          <v-card-title
            >Order {{ details.id }} | {{ details.name }}</v-card-title
          >
          <section>
            <v-card-text class="text-h7">
              Car {{ details.type }} | {{ details.brand }} <br />
              Problems {{ details.problem }} <br />
              {{ details.location }} km
              <v-row justify="end">
                <v-card-action>
                  <v-btn class="mx-1" variant="tonal" color="error">
                    Decline
                  </v-btn>

                  <v-btn class="mx-1" color="blue" variant="tonal">
                    Accept
                  </v-btn>
                </v-card-action>
              </v-row>
            </v-card-text>
          </section>
        </v-card>

        <v-dialog v-model="dialog" width="800" class="text-left pa-4">
          <v-card class="text-left pa-4">
            <v-row class="pa-5">
              <v-btn
                icon
                dark
                @click="dialog = false"
              >
              <v-icon>mdi-close</v-icon>
              </v-btn>
              <v-card-title class="text-h5"> Order {{details.id}}</v-card-title>
            </v-row>
            <v-card class="text-left ma-2 pa-1" variant="outlined">
              <v-title class="text-h6">Car Details</v-title>
              <v-card-text class="text-h7">
                Type: {{details.type}} | Brand: {{details.brand}}
                <br />
                License plate: {{details.plate}}
              </v-card-text>
            </v-card>
            <v-card
              variant="outlined"
              class="text-left ma-2 pa-1 d-flex justify-left align-center"
            >
              <section>
                <v-card-title>Current Location</v-card-title>
                <v-card-text>{{location.lat+', '+ location.lng}}</v-card-text>
              </section>
            </v-card>
            <v-card class="text-left ma-2 pa-1" variant="outlined">
              <v-title class="text-h6">Problems</v-title>
              <v-card-text class="text-h7">
                <p v-for="p in problems" :key="p">{{'-' + p}}</p>
              </v-card-text>
            </v-card>
            <v-card class="text-left ma-2 pa-1" variant="outlined">
              <v-title class="text-h6">Description</v-title>
              <v-card-text class="text-h7">
                <p>{{description}}</p>
              </v-card-text>
            </v-card> 
            <v-row justify="end" class="pa-4">
                <v-card-action>
                  <v-btn class="mx-1" variant="tonal" color="error" @click="dialog = false">
                    Decline
                  </v-btn>

                  <v-btn class="mx-1" color="blue" variant="tonal" @click="dialog = false">
                    Accept
                  </v-btn>
                </v-card-action>
              </v-row>
          </v-card>
        </v-dialog>

        <v-tabs v-model="tab" grow class="ma-4">
          <v-tab class="text-h5" value="one">On process</v-tab>
          <v-tab class="text-h5" value="two">Finish</v-tab>
        </v-tabs>

        <v-card-text>
          <v-window v-model="tab">
            <v-window-item value="one">
              <v-content>
                <v-container v-for="item in items" :key="item.id" class="pa-0 my-4">
                  <v-card width="100%" variant="tonal" class="text-left pa-4 ">
                    <v-row class="pa-2">
                      <v-card-title>
                        Order {{ item.id }} | {{ item.username }}
                      </v-card-title>
                      <v-chip class="ma-2" color="yellow-darken-3">{{
                        item.status
                      }}</v-chip>
                    </v-row>
                    <v-card-text class="text-h7">
                      Car {{ items.type }} | {{ item.brand }} <br />
                      Problems {{ item.problem }} <br />
                      {{ item.distance }} km
                    </v-card-text>
                  </v-card>
                </v-container>
              </v-content>
            </v-window-item>

            <v-window-item value="two">
              <v-card variant="tonal" class="text-left my-4 pa-4">
                <v-row class="pa-2">
                  <v-card-title
                    >Order {{ details.id }} | {{ details.name }}</v-card-title
                  >
                  <v-chip class="ma-2" color="green">{{details.status}}</v-chip>
                </v-row>
                <section>
                  <v-card-text class="text-h7">
                    Car {{ details.type }} | {{ details.brand }} <br />
                    Problems {{ details.problem }} <br />
                    {{ details.location }} km
                  </v-card-text>
                </section>
              </v-card>
            </v-window-item>
          </v-window>
        </v-card-text>
      </v-col>
      <v-col cols="4">
        <div class="mr-8">
        <v-container class="bg-grey-lighten-3 ma-4 rounded" style="height: 90vh;">
          <v-card class="mb-4" variant="tonal">
            <v-card-title class="text-left text-h5 pa-4">Total order</v-card-title>
              <v-card-text class="justify-center text-h4 pa-4">
                {{total}}
              </v-card-text>
              <v-card-text class="text-right text-h6 pa-4">
                Yesterday {{total}}
              </v-card-text>
          </v-card>
          <v-card variant="tonal" class="mb-4">
            <v-list-item class="text-left text-h6">Accepted
            <template v-slot:append>
              <v-badge
                color="green-lighten-1"
                :content="accept"
                :value="accept"
                inline
              ></v-badge>
            </template>
            </v-list-item>
          </v-card>
          <v-card variant="tonal" class="mb-4">
            <v-list-item class="text-left text-h6 pa-5">On the way
            <template v-slot:append>
              <v-badge
                color="green-lighten-1"
                :content="otw"
                :value="otw"
                inline
              ></v-badge>
            </template>
            </v-list-item>
          </v-card>
          <v-card variant="tonal" class="mb-4">
            <v-list-item class="text-left text-h6 pa-5">On process
            <template v-slot:append>
              <v-badge
                color="green-lighten-1"
                :content="onprocess"
                :value="onprocess"
                inline
              ></v-badge>
            </template>
            </v-list-item>
          </v-card>
          <v-card variant="tonal" class="mb-4"> 
            <v-list-item class="text-left text-h6 pa-5">Finish
            <template v-slot:append>
              <v-badge
                color="green-lighten-1"
                :content="finish"
                :value="finish"
                inline
              ></v-badge>
            </template>
            </v-list-item>
          </v-card>
        </v-container>
      </div>
      </v-col>
    </v-row>
  </div>
</template>

<style>
.text-h4 {
  display: flex;
  justify-content: left;
  align-content: center;
}
</style>

<script>
export default {
  data() {
    return {
      dialog: false,
      orders: 1,
      total: 5,
      accept: 2,
      otw: 1,
      onprocess: 1,
      finish: 1,
      problems: ["no battery", "broken motor"],
      description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam rutrum tincidunt arcu sed gravida. Suspendisse mollis ex sed magna viverra, eu tincidunt velit lobortis. Phasellus accumsan mauris est, in pretium orci lacinia in. Nulla in bibendum ex, eu condimentum mauris. Pellentesque quis nisl a justo ultrices molestie ac mattis libero. Aliquam vel sollicitudin quam, vel molestie nunc. Proin dolor dolor, vehicula sed nisl dapibus, semper scelerisque nisl. Sed id neque ac metus efficitur vestibulum. Phasellus quis nibh ac dui molestie fringilla sit amet bibendum sem. Quisque consectetur sit amet nunc ac elementum. Suspendisse vulputate mauris erat, nec tempus dui vulputate.",
      location: {
        lat: "20",
        lng: "90"
      },
      details: {
        id: "12345",
        type: "SUV",
        brand: "MG",
        plate: "ศง 4727",
        name: "nunnapat",
        problem: "2",
        location: "2.1",
        status: "Finish",
      },
      items: [
        {
          id: "11111",
          username: "nunnapat",
          distance: "2.2",
          type: "SUV",
          brand: "MG",
          problem: "2",
          status: "On-process",
        },
        {
          id: "22222",
          username: "natawat",
          distance: "0.4",
          type: "Sedan",
          brand: "MG",
          problem: "2",
          status: "On the way",
        },
        {
          id: "33333",
          username: "kiitiphum",
          distance: "1.3",
          type: "Van",
          brand: "MG",
          problem: "2",
          status: "Accept",
        },
      ],
      tab: null,
    };
  },
};
</script>