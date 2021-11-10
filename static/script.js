$(document).ready(function() {
    $("#go-button").click(function () {
        $("#cards").empty();
        const selectedOption = $("#select-options").find(':selected').val();
        $.get(`/foods/?option=${selectedOption}`, function(data) {
            data.forEach(element => {
                const card = `
                    <div class="card mt-1 w-50">
                        <div class="m-3">
                            <img src="${element.image}" class="card-img-top">
                        <div>
                        <div class="card-body">
                            <h5 class="card-title">${element.title}</h5>
                            <a href="${element.sourceUrl}" class="btn btn-primary" target="_blank">View Recipe</a>
                        </div>
                    </div>`;
                $("#cards").append(card);
            }); 
        });
    });
});