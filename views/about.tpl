



<!-- Main Content -->
<div class="container">
	<div class="row">
<!-- USE SIDEBAR -->
    <!-- PostList Container -->
    		<div class="col-lg-8 col-lg-offset-1
                col-md-8 col-md-offset-1
                col-sm-12
                col-xs-12
                postlist-container">

                <div class="zh post-container">
                    {{str2html .page.Body}}
                </div>
    		</div>

        <div class="
                col-lg-3 col-lg-offset-0
                col-md-3 col-md-offset-0
                col-sm-12
                col-xs-12
                sidebar-container
            ">
            <!-- Featured Tags -->

        {{template "layout/_includes/featured-tags.html" .}}


            <!-- Short About -->

        {{template "layout/_includes/short-about.html" .}}
            <!-- Friends Blog -->

        {{template "layout/_includes/friends.html" .}}
        </div>
	</div>
    <!-- Sidebar Container -->



</div>


