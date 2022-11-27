<template>
  <div class="text-center ma-4 justify-center">
    <div class="text-left text-h4 my-2">Order</div>
    <v-card>
      <v-tabs v-model="tab" center-active show-arrows bg-color="teal-darken-1">
        <v-tab class="text-subtitle-1" value="one">Incoming</v-tab>
        <v-tab class="text-subtitle-1" value="two">Accept</v-tab>
        <v-tab class="text-subtitle-1" value="three">On the way</v-tab>
        <v-tab class="text-subtitle-1" value="four">On process</v-tab>
        <v-tab class="text-subtitle-1" value="five">Finish</v-tab>
      </v-tabs>
    </v-card>

    <v-combobox
      v-model="chips"
      :items="filter"
      chips
      clearable
      label="Filter order"
      multiple
      prepend-icon="mdi-filter-variant"
      variant="solo"
      class="mt-4"
    >
      <template v-slot:selection="{ attrs, item, select, selected }">
        <v-chip
          v-bind="attrs"
          :model-value="selected"
          closable
          @click="select"
          @click:close="remove(item)"
        >
          <strong>{{ item }}</strong>&nbsp;
          <span>(interest)</span>
        </v-chip>
      </template>
    </v-combobox>

    <v-card-text class="pa-0 ma-0">
      <v-window v-model="tab">
        <v-window-item value="one">
          <v-content class="justify-start pa-0">
            <v-container v-for="item in items" :key="item.ticketID" class="pa-0">
              <v-card @click="dialog=true; acceptTicket(item)" width="100%" variant="tonal" v-if="show" class="text-left pa-0 my-4">
                <v-card-title class="text-h6">
                  Order {{item.ticketID}} | {{item.username}}
                </v-card-title>
                <v-card-subtitle>{{item.date}}</v-card-subtitle>
                <v-card-text class="text-h7">
                  Car {{ item.type }} | {{ item.brand }} <br />
                  Problems {{ item.problem }} <br />
                  {{ item.distance }} km
                </v-card-text>
                <v-row justify="end" class="pa-6">
                  <v-card-action>
                    <v-btn @click.stop="removeMessage()" class="mx-1" variant="tonal" color="error">
                      Decline
                    </v-btn>
                    <v-btn @click.stop="removeMessage()" class="mx-1" color="blue" variant="tonal">
                      Accept
                    </v-btn>
                  </v-card-action>
                </v-row>
              </v-card>
            </v-container>
          </v-content>
        </v-window-item>

            <v-window-item value="two">
              <v-content>
                <v-container v-for="item in accept" :key="item.id" class="pa-0 my-4">
                  <v-card :to="'/details'" variant="tonal" class="text-left my-4">
                    <v-row class="pa-2">
                      <v-card-title class="text-h6">
                        Order {{item.id}} | {{item.username}}
                      </v-card-title>
                      <v-chip class="ma-2" color="yellow-darken-3">{{item.status}}</v-chip>
                    </v-row>      
                    <v-card-subtitle>{{item.date}}</v-card-subtitle>
                    <v-card-text class="text-h7">
                      Car {{ item.type }} | {{ item.brand }} <br />
                      Problems {{ item.problem }} <br />
                      {{ item.distance }} km
                    </v-card-text>
                  </v-card>
                </v-container>
              </v-content>
            </v-window-item>
            
            <v-window-item value="three">
              <v-content>
                <v-container v-for="item in otw" :key="item.id" class="pa-0 my-4">
                  <v-card :to="'/details'" variant="tonal" class="text-left my-4">
                    <v-row class="pa-2">
                      <v-card-title class="text-h6">
                        Order {{item.id}} | {{item.username}}
                      </v-card-title>
                      <v-chip class="ma-2" color="yellow-darken-3">{{item.status}}</v-chip>
                    </v-row>      
                    <v-card-subtitle>{{item.date}}</v-card-subtitle>
                    <v-card-text class="text-h7">
                      Car {{ item.type }} | {{ item.brand }} <br />
                      Problems {{ item.problem }} <br />
                      {{ item.distance }} km
                    </v-card-text>
                  </v-card>
                </v-container>
              </v-content>
            </v-window-item>

            <v-window-item value="four">
              <v-content>
                <v-container v-for="item in process" :key="item.id" class="pa-0 my-4">
                  <v-card :to="'/details'" variant="tonal" class="text-left my-4">
                    <v-row class="pa-2">
                      <v-card-title class="text-h6">
                        Order {{item.id}} | {{item.username}}
                      </v-card-title>
                      <v-chip class="ma-2" color="yellow-darken-3">{{item.status}}</v-chip>
                    </v-row>      
                    <v-card-subtitle>{{item.date}}</v-card-subtitle>
                    <v-card-text class="text-h7">
                      Car {{ item.type }} | {{ item.brand }} <br />
                      Problems {{ item.problem }} <br />
                      {{ item.distance }} km
                    </v-card-text>
                  </v-card>
                </v-container>
              </v-content>
            </v-window-item>

            <v-window-item value="five">
              <v-content>
                <v-container v-for="item in finish" :key="item.id" class="pa-0 my-4">
                  <v-card :to="'/details'" variant="tonal" class="text-left my-4">
                    <v-row class="pa-2">
                      <v-card-title class="text-h6">
                        Order {{item.id}} | {{item.username}}
                      </v-card-title>
                      <v-chip class="ma-2" color="green">{{item.status}}</v-chip>
                    </v-row>      
                    <v-card-subtitle>{{item.date}}</v-card-subtitle>
                    <v-card-text class="text-h7">
                      Car {{ item.type }} | {{ item.brand }} <br />
                      Problems {{ item.problem }} <br />
                      {{ item.distance }} km
                    </v-card-text>
                  </v-card>
                </v-container>
              </v-content>
            </v-window-item>
          </v-window>
        </v-card-text>

        <v-dialog v-model="dialog2">
          <v-card>
            <v-card-title>Heelo</v-card-title>
          </v-card>
        </v-dialog>

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
          </v-card>
        </v-dialog>
  </div>
</template>

<style>

</style>

<script>
export default {
  data() {
    return {
      items: [
        {
          id: "11111",
          username: "nunnapat",
          date: "12 Nov 2022",
          distance: "2.2",
          type: "SUV",
          brand: "MG",

        },
        {
          id: "22222",
          username: "natawat",
          date: "18 Nov 2022",
          distance: "0.4",
          type: "Sedan",
          brand: "MG",
          problem: "2",
        },
        {
          id: "33333",
          username: "kiitiphum",
          date: "22 Nov 2022",
          distance: "1.3",
          type: "Van",
          brand: "MG",
        },
      ],
      accept: [
        {
          id: "11111",
          username: "nunnapat",
          date: "12 Nov 2022",
          distance: "2.2",
          type: "SUV",
          brand: "MG",
          problem: "2",
          status: "Accept",
        },
        {
          id: "22222",
          username: "natawat",
          date: "18 Nov 2022",
          distance: "0.4",
          type: "Sedan",
          brand: "MG",
          problem: "2",
          status: "Accept",
        },
        {
          id: "33333",
          username: "kiitiphum",
          date: "22 Nov 2022",
          distance: "1.3",
          type: "Van",
          brand: "MG",
          status: "Accept", 
        },
      ],
      otw: [
        {
          id: "11111",
          username: "nunnapat",
          date: "12 Nov 2022",
          distance: "2.2",
          type: "SUV",
          brand: "MG",
          problem: "2",
          status: "On the way",
        },
        {
          id: "22222",
          username: "natawat",
          date: "18 Nov 2022",
          distance: "0.4",
          type: "Sedan",
          brand: "MG",
          problem: "2",
          status: "On the way",
        },
        {
          id: "33333",
          username: "kiitiphum",
          date: "22 Nov 2022",
          distance: "1.3",
          type: "Van",
          brand: "MG",
          status: "On the way",
        },
      ],
      process: [
        {
          id: "11111",
          username: "nunnapat",
          date: "12 Nov 2022",
          distance: "2.2",
          type: "SUV",
          brand: "MG",
          problem: "2",
          status: "In process",
        },
        {
          id: "22222",
          username: "natawat",
          date: "18 Nov 2022",
          distance: "0.4",
          type: "Sedan",
          brand: "MG",
          problem: "2",
          status: "In process",
        },
        {
          id: "33333",
          username: "kiitiphum",
          date: "22 Nov 2022",
          distance: "1.3",
          type: "Van",
          brand: "MG",
          status: "In process",
        },
      ],
      finish: [
        {
          id: "11111",
          username: "nunnapat",
          date: "12 Nov 2022",
          distance: "2.2",
          type: "SUV",
          brand: "MG",
          problem: "2",
          status: "Finish",
        },
        {
          id: "22222",
          username: "natawat",
          date: "18 Nov 2022",
          distance: "0.4",
          type: "Sedan",
          brand: "MG",
          problem: "2",
          status: "Finish",
        },
        {
          id: "33333",
          username: "kiitiphum",
          date: "22 Nov 2022",
          distance: "1.3",
          type: "Van",
          brand: "MG",
          status: "Finish",
        },
      ],
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
      tab: null,
      dialog: false,
      dialog2: false,
      location: {
        lat: "20",
        lng: "90"
      },
      chips: [ ],
      show: true,
      filter: ['Date', 'Order no', 'A-Z', 'Z-A'],
      problems: ["no battery", "broken motor"],
      description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam rutrum tincidunt arcu sed gravida. Suspendisse mollis ex sed magna viverra, eu tincidunt velit lobortis. Phasellus accumsan mauris est, in pretium orci lacinia in. Nulla in bibendum ex, eu condimentum mauris. Pellentesque quis nisl a justo ultrices molestie ac mattis libero. Aliquam vel sollicitudin quam, vel molestie nunc. Proin dolor dolor, vehicula sed nisl dapibus, semper scelerisque nisl. Sed id neque ac metus efficitur vestibulum. Phasellus quis nibh ac dui molestie fringilla sit amet bibendum sem. Quisque consectetur sit amet nunc ac elementum. Suspendisse vulputate mauris erat, nec tempus dui vulputate.",
    };
  },
  methods:{
    remove (item) {
      this.chips.splice(this.chips.indexOf(item), 1)
      this.chips = [...this.chips]
    },
    removeMessage(seconds) {
         setTimeout(()=> this.show = false, seconds * 1000);
      },
  }
};
</script>