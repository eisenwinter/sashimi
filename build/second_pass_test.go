package build

import "github.com/eisenwinter/sashimi/resolver"

type mockNodeResolver struct{}

func (m *mockNodeResolver) ResolveByID(nodeType string, id int) (map[string]interface{}, error) {
	panic("nan")
}

func (m *mockNodeResolver) ResolveAll(nodeType string) ([]map[string]interface{}, error) {
	panic("nan")
}

func (m *mockNodeResolver) Resolve(nodeType string, predicate resolver.Predicate) ([]map[string]interface{}, error) {
	panic("nan")
}

var testDummyProject = map[string]string{
	"page.html": `<!doctype html><html><body><!-- sashimi:layout(main) -->
	<!--
	  sashimi:entity(page) of
	   - title as "Pagetitle" is text
	   - description as "Description" is text[:markdown]
	-->
	<!-- sashimi:display(page.description) -->
	<div class="md-content">
	  This content wil be replaced with the page contents et al
	</div>`,
	"projects.html": `<html><body>
	<!-- sashimi:layout(main) -->
	<!-- sashimi:repeat(project) as p -->
	<div class="project-item">
	   <!-- sashimi:display(p.title) -->
	   <h1>MyProject </h1>
	   <!-- sashimi:link(p) -->
	   <a href="#"></a>
	</div></body></html>`,
	"project.html": `<!-- 
	sashimi:entity(project) of
	 - title as "Project Title" is text
	 - description as "Description" is text
	 - image as "Main Image" is picture
	 - starCount as "Star Count" is number
	 - category as "Category" is oneOf("Archtiecture","Misc")
	 - tags as "Tags" is manyOf("Cheap","Fast","Good")
	 - gallery as "Gallery" is list picture
  -->
  <!doctype html>
  <html><body>
  <div>
    <!-- sashimi:display(project.title) -->
	<h1>MyProject </h1>
	<!-- sashimi:display(project.description) --> 
	<p>
	  Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusantium, placeat praesentium, cum tempore in tenetur quo voluptate eum ipsa commodi laudantium assumenda vero modi itaque sunt vel qui est ut.
	</p>
  </div></body></html>`,
	"shared/_layout.html": `<!doctype html>
	<html class="no-js" lang="">
	<head>
	  <meta charset="utf-8">
	  <title></title>
	  <meta name="description" content="">
	  <meta name="viewport" content="width=device-width, initial-scale=1">
	  <meta property="og:title" content="">
	  <meta property="og:type" content="">
	  <meta property="og:url" content="">
	  <meta property="og:image" content="">
	</head>
	<body>
	  <nav>
	   <ul>
		<!-- sashimi:repeat(@pages) as e -->
		<li>
		   <!-- sashimi:link(e) -->
		  <a href="#">Blah</a>
		</li>
	   </ul>
	  </nav>
	  <div class="main">
	  </div>
	  <!-- sashimi:layout_section(main) -->
	  <div class="content">
	  </div>
	</body>
	</html>
	`}
