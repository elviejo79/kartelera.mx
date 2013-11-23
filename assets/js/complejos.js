$(document).ready(function(){
		$( ".date" ).each(function() {
		$(this).click(function(){
			var number = $(this).attr("id").slice(7);

			var item = $("#complejos" + number);
			$(".complejos:visible").toggle();
			item.toggle();
		});
	});

});