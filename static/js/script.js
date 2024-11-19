let pieces = [];

const piecesTable = document.querySelector("#pieces-table tbody");
const pieceModal = document.getElementById("piece-modal");
const closeModal = document.getElementById("close-modal");
const pieceForm = document.getElementById("piece-form");
const submitBtn = document.getElementById("submit-btn");
const addPieceBtn = document.getElementById("add-piece-btn");

const renderPieces = () => {
    piecesTable.innerHTML = "";
    pieces.forEach(piece => {
        const row = document.createElement("tr");

        row.innerHTML = `
            <td>${piece.id}</td>
            <td>${piece.title}</td>
            <td>${piece.slug}</td>
            <td>${piece.value}</td>
            <td>
                <button class="edit-btn" data-id="${piece.id}">Edit</button>
                <button class="delete-btn" data-id="${piece.id}">Delete</button>
            </td>
        `;
        
        piecesTable.appendChild(row);
    });
};

const openModal = (piece = null) => {
    pieceModal.style.display = "block";

    if (piece) {
        document.getElementById("modal-title").textContent = "Edit Art Piece";
        document.getElementById("title").value = piece.title;
        document.getElementById("slug").value = piece.slug;
        document.getElementById("value").value = piece.value;
        document.getElementById("description").value = piece.description;
        document.getElementById("details").value = piece.details;
        submitBtn.textContent = "Update";
        pieceForm.onsubmit = (e) => handleUpdate(e, piece.id);
    } else {
        document.getElementById("modal-title").textContent = "Add New Art Piece";
        pieceForm.reset();
        submitBtn.textContent = "Add Piece";
        pieceForm.onsubmit = handleAdd;
    }
};

const closeModalFunction = () => {
    pieceModal.style.display = "none";
};

const handleAdd = async (e) => {
    e.preventDefault();

    const newPiece = {
        title: document.getElementById("title").value,
        slug: document.getElementById("slug").value,
        value: parseFloat(document.getElementById("value").value),
        description: document.getElementById("description").value,
        details: document.getElementById("details").value,
    };

    try {
        const response = await fetch("/api/pieces", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(newPiece),
        });

        if (response.ok) {
            const createdPiece = await response.json();
            pieces.push(createdPiece);
            renderPieces();
            closeModalFunction();
        } else {
            console.error("Failed to add piece");
        }
    } catch (error) {
        console.error("Error:", error);
    }
};

const handleUpdate = async (e, id) => {
    e.preventDefault();

    const updatedPiece = {
        title: document.getElementById("title").value,
        slug: document.getElementById("slug").value,
        value: parseFloat(document.getElementById("value").value),
        description: document.getElementById("description").value,
        details: document.getElementById("details").value,
    };

    try {
        const response = await fetch(`/api/pieces/${id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(updatedPiece),
        });

        if (response.ok) {
            const updatedPieceData = await response.json();
            pieces = pieces.map(piece => piece.id === id ? updatedPieceData : piece);
            renderPieces();
            closeModalFunction();
        } else {
            console.error("Failed to update piece");
        }
    } catch (error) {
        console.error("Error:", error);
    }
};

const handleDelete = async (id) => {
    try {
        const response = await fetch(`/api/pieces/${id}`, {
            method: "DELETE",
        });

        if (response.ok) {
            pieces = pieces.filter(piece => piece.id !== id);
            renderPieces();
        } else {
            console.error("Failed to delete piece");
        }
    } catch (error) {
        console.error("Error:", error);
    }
};

const fetchPieces = async () => {
    try {
        const response = await fetch("/api/pieces");
        if (response.ok) {
            pieces = await response.json();
            console.log(pieces)
            renderPieces();
        } else {
            console.error("Failed to fetch pieces");
        }
    } catch (error) {
        console.error("Error:", error);
    }
};

addPieceBtn.addEventListener("click", () => openModal());
closeModal.addEventListener("click", closeModalFunction);
piecesTable.addEventListener("click", (e) => {
    if (e.target.classList.contains("edit-btn")) {
        const id = parseInt(e.target.getAttribute("data-id"));
        const piece = pieces.find(p => p.id === id);
        openModal(piece);
    }

    if (e.target.classList.contains("delete-btn")) {
        const id = parseInt(e.target.getAttribute("data-id"));
        handleDelete(id);
    }
});

fetchPieces();
