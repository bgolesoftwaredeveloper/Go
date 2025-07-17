// Example usage of the Treap data structure.
// This program demonstrates insertion, in-order traversal, search, and memory cleanup
// using the Treap implementation from the treap package.
//
// Author: Braiden Gole
// Created: July 17, 2025
package main

import (
	"fmt"

	bi_directional "github.com/bgolesoftwaredeveloper/bi_directional/BiDirectionalImplementation"
)

func main() {
	fmt.Println("Pharmaceutical bi-directional tree: [Desvenlafaxine]")
	fmt.Println()

	// Root node.
	var pharmaceutical *bi_directional.Node = &bi_directional.Node{Value: "Pharmaceutical"}

	var desvenlafaxine *bi_directional.Node = &bi_directional.Node{Value: "Desvenlafaxine"}

	pharmaceutical.AddChildNode(desvenlafaxine)

	// Pharmacokinetics subtree.
	var pharmacokinetics *bi_directional.Node = &bi_directional.Node{Value: "Pharmacokinetics"}

	pharmacokinetics.AddChild("Absorption: Rapidly and completely absorbed after oral administration")
	pharmacokinetics.AddChild("Distribution: Widely distributed, moderate volume of distribution")
	pharmacokinetics.AddChild("Metabolism: Metabolized primarily by conjugation")
	pharmacokinetics.AddChild("Elimination: Excreted mainly in urine as unconjugated and conjugated desvenlafaxine")

	// Add pharmacokinetics as a child (just the label, since AddChild takes string).
	desvenlafaxine.AddChildNode(pharmacokinetics)

	// Add key drug attributes/details directly as children of Desvenlafaxine node.
	pharmaceutical.AddChildNode(&bi_directional.Node{Value: "Active Ingredients: Desvenlafaxine succinate"})
	pharmaceutical.AddChildNode(&bi_directional.Node{Value: "Drug Class: Antidepressant, Serotonin-Norepinephrine Reuptake Inhibitor (SNRI)"})
	pharmaceutical.AddChildNode(&bi_directional.Node{Value: "Chemical Structure: C₁₆H₂₅NO₃"})
	pharmaceutical.AddChildNode(&bi_directional.Node{Value: "Mechanism of Action: Inhibition of serotonin and norepinephrine reuptake"})

	// Pharmacodynamics
	var pharmacodynamics *bi_directional.Node = &bi_directional.Node{Value: "Pharmacodynamics"}

	pharmacodynamics.AddChild("Increases availability of serotonin and norepinephrine in the synaptic cleft")
	pharmacodynamics.AddChild("Therapeutic effects observed within the first 1-2 weeks of treatment")

	desvenlafaxine.AddChildNode(pharmacodynamics)

	// Indications
	var indications *bi_directional.Node = &bi_directional.Node{Value: "Indications"}

	indications.AddChild("Major depressive disorder (MDD)")

	desvenlafaxine.AddChildNode(indications)

	// Contraindications
	var contraindications *bi_directional.Node = &bi_directional.Node{Value: "Contraindications"}

	contraindications.AddChild("Hypersensitivity to desvenlafaxine or any component of the formulation")
	contraindications.AddChild("Concomitant use of monoamine oxidase inhibitors (MAOIs)")

	desvenlafaxine.AddChildNode(contraindications)

	// Adverse effects
	var adverseEffects *bi_directional.Node = &bi_directional.Node{Value: "Adverse Effects"}

	adverseEffects.AddChild("Nausea, dry mouth, constipation, decreased appetite, increased blood pressure")
	adverseEffects.AddChild("Insomnia, dizziness, sweating, sexual dysfunction")

	desvenlafaxine.AddChildNode(adverseEffects)

	// Drug interactions
	var drugInteractions *bi_directional.Node = &bi_directional.Node{Value: "Drug Interactions"}

	drugInteractions.AddChild("Concomitant use of MAOIs, other serotonergic agents, CYP3A4 inhibitors/inducers")

	desvenlafaxine.AddChildNode(drugInteractions)

	// Dosage forms
	var dosageForms *bi_directional.Node = &bi_directional.Node{Value: "Dosage Forms"}

	dosageForms.AddChild("Oral extended-release tablets")

	desvenlafaxine.AddChildNode(dosageForms)

	// Regulatory status
	desvenlafaxine.AddChild("Regulatory Status: FDA approved for the treatment of major depressive disorder")

	// Print the entire tree pharmaceutical.
	desvenlafaxine.PrintDown(0)

	fmt.Println()
}
