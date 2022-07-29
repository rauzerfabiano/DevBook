$('#parar-de-seguir').on('click', pararDeSeguir)
$('#seguir').on('click', seguir)
$('#editar-usuario').on('submit', editar)

function pararDeSeguir() {
    const usuarioId = $(this).data('usuario-id');

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: 'POST',
    }).done(function(){
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao parar de seguir o usuário!", "error");
    });
}

function seguir() {
    const usuarioId = $(this).data('usuario-id');

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: 'POST',
    }).done(function(){
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao seguir o usuário!", "error");
    });
}

function editar(evento){
    evento.preventDefault();

    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
        }
    }).done(function(){
        Swal.fire("Sucesso!", "Usuário atualizado com sucesso!", "success")
        .then(function(){
            window.location = "/perfil";
        })
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao atualizar o usuário!", "error");
    })
}