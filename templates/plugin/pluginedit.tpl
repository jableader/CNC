<md-content class="content" flex ng-show="main.focus == 'plugin-edit'">
  <md-data-table-toolbar>
    <h2 class="md-title" flex="50" ng-click="main.activateRouted('/admin/plugins', 'plugins');">
      <md-icon md-font-library="material-icons">keyboard_arrow_left</md-icon>
      Plugins <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon> Edit</h2>

    <div class="md-toolbar-tools" md-hide="showLoading" >
        <md-progress-linear md-mode="{{loadingMode()}}"></md-progress-linear>
    </div>

  </md-data-table-toolbar>

  <style>
  .small-icons {
    min-width: 16px;
    padding: 2px;
    margin: 2px;
  }
  </style>



  <md-content flex layout="row" layout-fill layout-wrap>
    <md-content md-margin md-padding flex="99" flex-sm="100" layout="column" layout-fill >

      <md-input-container flex layout-fill>
        <label>Name</label>
        <input ng-model="plugin.Name" type="text">
      </md-input-container>

      <!-- TODO: Add icon when https://github.com/angular/material/issues/6150 is finished -->
      <md-autocomplete flex required
        id="iconSearchText"
        md-input-name="autocompleteField"
        md-items="icon in getIcons()"
        md-search-text="iconSearchText"
        md-min-length="0"
        md-item-text="icon"
        md-selected-item="plugin.Icon"
        md-floating-label="Icon">
          <md-item-template>
            <md-icon md-font-library="material-icons" md-highlight-text="icon">{{icon}}</md-icon>
            {{icon}}
          </md-item-template>
      </md-autocomplete>

      <md-input-container class="md-block">
            <label>Description</label>
            <textarea ng-model="plugin.Description" columns="1" md-maxlength="150" rows="5"></textarea>
      </md-input-container>

      <md-switch ng-model="plugin.Enabled" aria-label="Enabled" class="md-block">
        Enabled
      </md-switch>
    </md-content>

    <md-content md-margin md-padding flex="42" flex-sm="100" layout="column">
      <p style="color: rgba(0, 0, 0, 0.54);">Resources
        <md-button class="ng-icon-button" ng-click="createNewResource()" aria-label="Add Resource">
          <md-icon md-font-library="material-icons">add</md-icon>
        </md-button>
      </p>
      <md-data-table-container ng-hide="showLoading || (plugin.Resources.length == 0)">
        <table md-data-table md-progress="deferred">
          <thead>
            <tr>
              <th name="Name"></th>
              <th name="Actions"></th>
            </tr>
          </thead>
          <tbody>
            <tr ng-repeat="resource in plugin.Resources">
              <td>
                <md-icon ng-show="resource.ResType == 'JSC'" md-font-library="material-icons">code</md-icon>
                <md-icon ng-show="resource.ResType == 'TPL'" md-font-library="material-icons">description</md-icon>
                 {{resource.Name}}</td>
              <td>
                <md-button class="ng-icon-button small-icons" ng-click="main.activateRouted('/admin/resource/'+resource.ID, 'resource-form')" aria-label="Edit Resource">
                  <md-icon md-font-library="material-icons">mode_edit</md-icon>
                </md-button>
                <md-button class="ng-icon-button small-icons" ng-click="deleteResource(resource.ID, $event)" aria-label="Delete Resource">
                  <md-icon md-font-library="material-icons">delete</md-icon>
                </md-button>
              </td>
            </tr>
          </tbody>
        </table>
      </md-data-table-container>
    </md-content>
  </md-content>

  <button flex="" layout-fill=""
  class="md-raised md-primary md-button md-scope"
  ng-click="process()"
  aria-label="Save Plugin" tabindex="0" aria-disabled="true">
   <i class="material-icons" style="vertical-align: middle;">save</i>
  <span style="vertical-align: middle;"> Commit Changes</span>
  </button>

</md-content>
