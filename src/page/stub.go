package page

import (
	_ "embed"
)

// func WorkflowsSearchEndpoint(ctx *gin.Context) {
// 	searchTerm, ok := ctx.GetQuery("search")
// 	if !ok {
// 		ctx.Status(http.StatusBadRequest)
// 		return
// 	}
// 	query := strings.TrimSpace(searchTerm)

// 	stubs := GetStubs()

// 	filtered := []workflow.Workflow{}

// 	for name, ws := range stubs {
// 		if strings.Contains(strings.ToLower(name), strings.ToLower(query)) {
// 			filtered = append(filtered, s)
// 		}
// 	}

// 	markdown := stubBuildTable(filtered, ctx)

// 	ctx.Data(http.StatusOK, "text/html; charset=utf-8", markdown)
// }

// func stubBuildTable(stubs StubType, ctx *gin.Context) []byte {
// 	sort.Slice(stubs, func(i, j int) bool {
// 		return stubs[i].Name < stubs[j].Name
// 	})

// 	t := table.Table{
// 		ID: "stub_table",
// 		Headers: []header.Header{
// 			{
// 				Contents: "Name",
// 				Classes:  "text-lg",
// 			},
// 			{
// 				Contents: "Created",
// 				Classes:  "text-lg",
// 			},
// 			{
// 				Contents: "Updated",
// 				Classes:  "text-lg",
// 			},
// 			{
// 				Contents: "Version",
// 				Classes:  "text-lg",
// 			},
// 			{
// 				Contents: "",
// 				Classes:  "text-lg",
// 			},
// 			{
// 				Contents: "",
// 				Classes:  "text-lg",
// 			},
// 		},
// 		Rows:          make([][]cell.Cell, 0),
// 		Classes:       "theme-light",
// 		Style:         "width:100%;",
// 		HeaderClasses: "rounded-md",
// 	}

// 	for _, s := range stubs {
// 		r := []cell.Cell{
// 			{
// 				Contents: s.Name,
// 			},
// 			{
// 				Contents: s.Created,
// 			},
// 			{
// 				Contents: s.Updated,
// 			},
// 			{
// 				Contents: s.Version,
// 			},
// 			{
// 				Contents: `<a href="/ui/stubs/` + s.Name + `" class="table-link-link w3-right-align dark theme-text"
//                     style="float:right;margin-right:16px;">
//                     <i class="fa-solid fa-link"></i>
//                 </a>`,
// 			},
// 			{
// 				Contents: "",
// 			},
// 		}
// 		t.Rows = append(t.Rows, r)
// 	}

// 	html, err := t.Render()
// 	if err != nil {
// 		logger.Errorf("", "Cannot render stub table: %s", err.Error())
// 		ctx.AbortWithStatus(http.StatusInternalServerError)
// 		return []byte{}
// 	}
// 	return []byte(html)
// }

// func StubPageEndpoint(ctx *gin.Context) {
// 	markdown := stubBuildPage(ctx)
// 	ctx.Data(http.StatusOK, "text/html; charset=utf-8", markdown)
// }

// func stubBuildPage(ctx *gin.Context) []byte {
// 	p := page.Page{
// 		ID:             "page",
// 		SidebarEnabled: true,
// 		Sidebar: sidebar.Sidebar{
// 			ID:         "sidebar",
// 			Classes:    "theme-light",
// 			Components: []ui.Component{
// 				// Add links here
// 			},
// 		},
// 		Components: []ui.Component{
// 			// Add components here
// 		},
// 	}
// 	html, err := p.Render()
// 	if err != nil {
// 		logger.Errorf("", "Cannot render stub page: %s", err.Error())
// 		ctx.AbortWithStatus(http.StatusInternalServerError)
// 		return []byte{}
// 	}
// 	return []byte(html)
// }
