<!DOCTYPE html>
<html lang="en">

  <head>
      <title>CNC Plugin API Reference</title>
       <link href="/static/css/v-accordion.min.css" rel="stylesheet">
      {!{template "headcontent"}!}
  </head>

  <body layout="column"  ng-app="apiApp" ng-cloak>
      <md-toolbar layout="row" flex="10">
        <h1 flex><md-icon md-font-library="material-icons" style="font-size: 250%;">code</md-icon> CNC Plugin API Reference</h1>
      </md-toolbar>

    <div layout="row" flex>
      <div flex layout="column" tabIndex="-1" role="main" class="md-whiteframe-z2">

        <md-content md-padding md-margin flex layout-fill style="padding-left: 12px; padding-right: 12px;">
          <p>This document outlines the available features within plugins.</p>

          <v-accordion class="vAccordion--default" multiple>

            <!-- LOGGING FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">list</md-icon>
                Logging
              </v-pane-header>

              <v-pane-content>
                <p>These functions allow you to write to the system log and to stdout of the server process.</p>

                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      log(
                      <i class="amber">message</i> <sup style="color: #444444;">str</sup> )
                    </v-pane-header>
                    <v-pane-content>
                      Posts an INFO message to the system log (available in the summary tab). The message is tagged with
                      the name of the plugin prefixed with 'PLUGIN-'.
                      <p>EG: </p>
                      <pre>
                      log("Your shit."); //Output: [I] [PLUGIN-pluginnamehere] Your Shit.
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      console.log(
                      <i class="amber">message</i> <sup style="color: #444444;">str</sup> ... )
                    </v-pane-header>
                    <v-pane-content>
                      Writes data out to standard output. Should not use unless your an idiot or you dont want to dirty the system log.
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>







            <!-- GMAIL FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">mail</md-icon>
                Gmail
              </v-pane-header>

              <v-pane-content>
                <p>These functions allow you to send emails from a Gmail address, given a username and password. If
                  Two-factor authentication is used on the account, you will need to generate an 'app specific password' (google it).</p>

                <p>There are two steps to sending an email. First, you should call <span class="green">setup()</span>
                to initialise the system with the correct credentials. Next, you should immediately call <span class="green">sendMessage()</span>
                with your subject, a string list of recipients, and the contents (optionally html) of your email.</p>

                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      gmail.setup(
                      <i class="amber">username</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">password</i> <sup style="color: #444444;">str</sup> )
                    </v-pane-header>
                    <v-pane-content>
                      Sets up the system with the correct credentials for when you call an email. NOTE call this immediately before you call
                      sendMessage().
                      <p>EG: </p>
                      <pre>
                      gmail.setup('jacob@gmail.com', 'pass');
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      gmail.sendMessage(
                      <i class="amber">subject</i> <sup style="color: #444444;">str</sup> ,
                      <i class="amber">recipients</i> <sup style="color: #444444;">list of strings</sup> ,
                      <i class="amber">content</i> <sup style="color: #444444;">str</sup> )
                    </v-pane-header>
                    <v-pane-content>
                      Sets up the system with the correct credentials for when you call an email. NOTE call this immediately before you call
                      sendMessage().
                      <p>EG: </p>
                      <pre>
                      gmail.sendMessage('Test Subject', ['barry@gmail.com'], 'KEK BRAAAH');
                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>




            <!-- CRON FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">build</md-icon>
                CRON
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to schedule a method to be run at a certain time, or ever x seconds/minutes/hours</p>

                <p>The cron format used when writing <i>schedule descriptor</i> is described in <a href="https://godoc.org/github.com/robfig/cron">https://godoc.org/github.com/robfig/cron</a></p>
                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      cron.schedule(
                      <i class="amber">schedule descriptor</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">method name</i> <sup style="color: #444444;">str</sup> )
                    </v-pane-header>
                    <v-pane-content>
                      Sets up a method to be called at the time specified by <i>schedule descriptor</i>.
                      <p>EG: </p>
                      <pre>
                        function periodicCron(dummy) {
                            log("CRON trigger activated")
                        }

                        cron.schedule("@every 15m", "periodicCron");
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header style="color: #000077;">
                      <md-icon md-font-library="material-icons">call_missed </md-icon>
                      function methodname (
                      <i class="amber">param1</i> <sup style="color: #444444;">obj</sup> ->
                      (
                        <i class="amber">CronID</i> <sup style="color: #444444;">str</sup>
                      )
                      ) { }
                    </v-pane-header>
                    <v-pane-content>
                      Describes the callback of the method whoes name is passed in cron.schedule().

                      <p>EG: </p>
                      <pre>
                        function periodicCron(dummy) {
                            log("CRON trigger activated")
                        }

                        cron.schedule("@every 15m", "periodicCron");
                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>






            <!-- TWILIO FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">message</md-icon>
                Twilio
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to send SMS'es. Be advised this costs the owner (about 6c per SMS).</p>

                <p>The owner will have one or more numbers setup from which can be SMS'ed. One needs to be known to send SMS'es.</p>
                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      twilio.sendSMS(
                      <i class="amber">from number</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">to number</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">message</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      To and from addresses must be in the format: +COUNTRYCODENUMBER.
                      <p>EG: </p>
                      <pre>
                        var fromNumber = "+19284332644";
                        var toAddress = "+61342320983";

                        twilio.sendSMS(fromNumber, toAddress, "YOLO!!! :D");

                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>













            <!-- WEB FUNCTIONS -->
            <v-pane disabled>
              <v-pane-header>
                <md-icon md-font-library="material-icons">cloud</md-icon>
                Web Handlers (Documentation coming soon)
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to send SMS'es. Be advised this costs the owner (about 6c per SMS).</p>

                <p>The owner will have one or more numbers setup from which can be SMS'ed. One needs to be known to send SMS'es.</p>
                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      twilio.sendSMS(
                      <i class="amber">from number</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">to number</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">message</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      To and from addresses must be in the format: +COUNTRYCODENUMBER.
                      <p>EG: </p>
                      <pre>
                        var fromNumber = "+19284332644";
                        var toAddress = "+61342320983";

                        twilio.sendSMS(fromNumber, toAddress, "YOLO!!! :D");

                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>





            <!-- Web Request FUNCTIONS -->
            <v-pane disabled>
              <v-pane-header>
                <md-icon md-font-library="material-icons">http</md-icon>
                Web Requests (Documentation coming soon)
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to send SMS'es. Be advised this costs the owner (about 6c per SMS).</p>

                <p>The owner will have one or more numbers setup from which can be SMS'ed. One needs to be known to send SMS'es.</p>
                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      twilio.sendSMS(
                      <i class="amber">from number</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">to number</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">message</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      To and from addresses must be in the format: +COUNTRYCODENUMBER.
                      <p>EG: </p>
                      <pre>
                        var fromNumber = "+19284332644";
                        var toAddress = "+61342320983";

                        twilio.sendSMS(fromNumber, toAddress, "YOLO!!! :D");

                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>

                </v-accordion>
        </md-content>
      </div>
    </div>

    {!{template "tailcontent"}!}

    <script>
    (function() {

        var app = angular.module('apiApp', [
          'md.data.table',
          'vAccordion',
          'ngMaterial']);

        //material colour scheme
        app.config(function($mdThemingProvider, $mdIconProvider){
          $mdThemingProvider.theme('default')
                              .primaryPalette('indigo')
                              .accentPalette('amber');
        });

    })();
    </script>

    <script src="/static/js/v-accordion.min.js"></script>
  </body>
</html>
